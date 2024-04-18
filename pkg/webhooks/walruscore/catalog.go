package walruscore

import (
	"context"
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/webhook"
)

// CatalogWebhook hooks a v1.Catalog object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="catalogs",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE","DELETE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type CatalogWebhook struct {
	webhook.DefaultCustomValidator
}

func (r *CatalogWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	return &walruscore.Catalog{}, nil
}

var _ ctrlwebhook.CustomValidator = (*CatalogWebhook)(nil)

func (r *CatalogWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	return r.validateFilter(obj)
}

func (r *CatalogWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	return r.validateFilter(newObj)
}

func (r *CatalogWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	// TODO(michelia): add validate to prevent deletion of catalog depended by resource definition/resource/resource run.
	return nil, nil
}

func (r *CatalogWebhook) validateFilter(obj runtime.Object) (ctrladmission.Warnings, error) {
	c := obj.(*walruscore.Catalog)

	filters := c.Spec.Filters
	if filters == nil {
		return nil, nil
	}

	if filters.IncludeExpression != "" {
		if _, err := regexp.Compile(filters.IncludeExpression); err != nil {
			return nil, field.Invalid(
				field.NewPath("spec.filters.includeExpression"), filters.IncludeExpression, err.Error())
		}
	}
	if filters.ExcludeExpression != "" {
		if _, err := regexp.Compile(filters.ExcludeExpression); err != nil {
			return nil, field.Invalid(
				field.NewPath("spec.filters.excludeExpression"), filters.ExcludeExpression, err.Error())
		}
	}
	return nil, nil
}
