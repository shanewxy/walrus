package walruscore

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/webhook"
)

// TemplateWebhook hooks a v1.Template object.
//
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="templates",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","DELETE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type TemplateWebhook struct {
	webhook.DefaultCustomValidator
}

func (r *TemplateWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	return &walruscore.Template{}, nil
}

var _ ctrlwebhook.CustomValidator = (*TemplateWebhook)(nil)

func (r *TemplateWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	t := obj.(*walruscore.Template)
	_, err := sourceurl.ParseURLToSourceURL(t.Spec.VCSRepository.URL)
	if err != nil {
		return nil, field.Invalid(
			field.NewPath("spec.vcsRepository.url"), t.Spec.VCSRepository.URL, err.Error())
	}
	return nil, nil
}

func (r *TemplateWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	// TODO(michelia): add validate to prevent deletion of template depended by resource definition/resource/resource run.
	return nil, nil
}
