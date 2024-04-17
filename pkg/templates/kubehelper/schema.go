package kubehelper

import (
	"bytes"
	"context"

	"github.com/seal-io/utils/json"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/templates/api"
)

const (
	SchemaUserEditedNote string = "schema-user-edited"
)

// GenTemplateVersion generates a template version with schema.
func GenTemplateVersion(
	ctx context.Context,
	url, version string,
	t *walruscore.Template,
	sg *api.SchemaGroup,
) (*walruscore.TemplateVersion, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	var (
		schemaName           = NormalizeTemplateVersionSchemaName(t.Name, version, walruscore.NameSuffixTemplateSchema)
		originalUISchemaName = NormalizeTemplateVersionSchemaName(t.Name, version, walruscore.NameSuffixOriginalUISchema)
		uiSchemaName         = NormalizeTemplateVersionSchemaName(t.Name, version, walruscore.NameSuffixUISchema)
	)

	sb, err := json.Marshal(sg.Template)
	if err != nil {
		return nil, err
	}

	osb, err := json.Marshal(sg.UI)
	if err != nil {
		return nil, err
	}

	uisb, err := getUserEditedUISchema(ctx, t.Namespace, uiSchemaName)
	if err != nil {
		return nil, err
	}

	if uisb == nil {
		uisb = osb
	}

	err = createOrUpdateSchema(ctx, loopbackKubeCli, t, schemaName, sb)
	if err != nil {
		return nil, err
	}

	err = createOrUpdateSchema(ctx, loopbackKubeCli, t, originalUISchemaName, osb)
	if err != nil {
		return nil, err
	}

	err = createOrUpdateSchema(ctx, loopbackKubeCli, t, uiSchemaName, uisb)
	if err != nil {
		return nil, err
	}

	// Generate template version.
	tv := &walruscore.TemplateVersion{
		Version:              version,
		URL:                  url,
		TemplateSchemaName:   &schemaName,
		OriginalUISchemaName: &originalUISchemaName,
		UISchemaName:         &uiSchemaName,
	}

	return tv, nil
}

func createOrUpdateSchema(ctx context.Context, loopbackKubeCli clientset.Interface, t *walruscore.Template, name string, data []byte) error {
	cli := loopbackKubeCli.WalruscoreV1().Schemas(t.Namespace)

	var es *walruscore.Schema
	{
		es = &walruscore.Schema{
			ObjectMeta: meta.ObjectMeta{
				Name:      name,
				Namespace: t.Namespace,
			},
			Status: walruscore.SchemaStatus{
				Value: runtime.RawExtension{
					Raw: data,
				},
				Project: systemmeta.GetProjectName(t.Namespace),
			},
		}
		kubemeta.ControlOn(es, t, walruscore.SchemeGroupVersion.WithKind("Template"))
	}

	alignFn := func(as *walruscore.Schema) (*walruscore.Schema, bool, error) {
		if bytes.Equal(as.Status.Value.Raw, data) {
			return as, true, nil
		}

		as.Status.Value = runtime.RawExtension{
			Raw: data,
		}
		return as, false, nil
	}

	_, err := kubeclientset.Update(ctx, cli, es,
		kubeclientset.WithUpdateAlign(alignFn),
		kubeclientset.WithCreateIfNotExisted[*walruscore.Schema]())
	return err
}

// getUserEditedUISchema get exist user edited ui schema.
func getUserEditedUISchema(ctx context.Context, namespace, name string) ([]byte, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	existed, err := loopbackKubeCli.WalruscoreV1().Schemas(namespace).Get(ctx, name, meta.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	if len(existed.Status.Value.Raw) == 0 {
		return nil, nil
	}

	if systemmeta.DescribeResourceNote(existed, SchemaUserEditedNote) == "true" {
		return existed.Status.Value.Raw, nil
	}

	return nil, nil
}
