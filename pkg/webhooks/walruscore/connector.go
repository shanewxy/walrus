package walruscore

import (
	"context"
	"fmt"
	"reflect"

	"github.com/seal-io/utils/stringx"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/resourcehandlers"
	"github.com/seal-io/walrus/pkg/webhook"
)

// ConnectorWebhook hooks a v1.Connector object.
//
// nolint: lll
// +k8s:webhook-gen:mutating:group="walruscore.seal.io",version="v1",resource="connectors",scope="Namespaced"
// +k8s:webhook-gen:mutating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="NoneOnDryRun",matchPolicy="Equivalent",timeoutSeconds=10
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="connectors",scope="Namespaced",subResources=["status"]
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE","DELETE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type ConnectorWebhook struct {
	client ctrlcli.Client
}

func (r *ConnectorWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	r.client = opts.Manager.GetClient()

	return &walruscore.Connector{}, nil
}

var _ ctrlwebhook.CustomValidator = (*ConnectorWebhook)(nil)

func (r *ConnectorWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	logger := ctrllog.FromContext(ctx)

	conn := obj.(*walruscore.Connector)
	if stringx.StringWidth(conn.Name) > 30 {
		err := r.deleteSecret(ctx, conn)
		if err != nil {
			logger.Error(err, "failed to delete secret")
		}

		return nil, field.TooLongMaxLength(field.NewPath("name"), stringx.StringWidth(conn.Name), 30)
	}

	err := resourcehandlers.IsConnected(ctx, conn, r.client)
	if err != nil {
		derr := r.deleteSecret(ctx, conn)
		if derr != nil {
			logger.Error(derr, "failed to delete secret")
		}

		return nil, err
	}

	return nil, nil
}

func (r *ConnectorWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	conn := newObj.(*walruscore.Connector)
	if !reflect.DeepEqual(oldObj.(*walruscore.Connector).Spec, conn.Spec) {
		err := resourcehandlers.IsConnected(ctx, conn, r.client)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (r *ConnectorWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	conn := obj.(*walruscore.Connector)
	labelSelector := labels.SelectorFromSet(labels.Set{
		"walrus.seal.io/connector": fmt.Sprintf("%s-%s", conn.Namespace, conn.Name),
	})

	cbList := new(walruscore.ConnectorBindingList)
	err := r.client.List(ctx, cbList, ctrlcli.MatchingLabelsSelector{Selector: labelSelector})
	if err != nil {
		return nil, err
	}

	if len(cbList.Items) > 0 {
		return nil, field.Forbidden(field.NewPath("metadata", "name"), fmt.Sprintf("used by environment %s", cbList.Items[0].Namespace))
	}

	return nil, nil
}

func (r *ConnectorWebhook) Default(ctx context.Context, obj runtime.Object) error {
	conn := obj.(*walruscore.Connector)

	if conn.DeletionTimestamp != nil {
		return nil
	}

	data := map[string][]byte{}
	for k, v := range conn.Spec.Config.Data {
		data[k] = []byte(v.Value)
	}

	name := fmt.Sprintf("connector-config-%s", stringx.SumByFNV64a(conn.Namespace, conn.Name))
	eSec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: conn.Namespace,
			Name:      name,
		},
		Data: data,
	}

	_, err := kubeclientset.UpdateWithCtrlClient(ctx, r.client, eSec, kubeclientset.WithCreateIfNotExisted[*core.Secret]())
	if err != nil {
		return err
	}

	configData := map[string]walruscore.ConnectorConfigEntry{}
	for k, v := range conn.Spec.Config.Data {
		if v.Visible {
			configData[k] = v
		} else {
			configData[k] = walruscore.ConnectorConfigEntry{
				Value:   "",
				Visible: false,
			}
		}
	}

	conn.Spec.Config.Data = configData
	conn.Spec.SecretName = name

	kubemeta.SanitizeLastAppliedAnnotation(conn)

	return nil
}

func (r *ConnectorWebhook) deleteSecret(ctx context.Context, conn *walruscore.Connector) error {
	name := conn.Spec.SecretName
	eSec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: conn.Namespace,
			Name:      name,
		},
	}

	return kubeclientset.DeleteWithCtrlClient(ctx, r.client, eSec)
}
