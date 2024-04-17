package walruscore

import (
	"context"
	"fmt"
	"reflect"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/webhook"
)

// ConnectorBindingWebhook hooks a v1.ConnectorBinding object.
//
// nolint: lll
// +k8s:webhook-gen:mutating:group="walruscore.seal.io",version="v1",resource="connectorbindings",scope="Namespaced"
// +k8s:webhook-gen:mutating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="NoneOnDryRun",matchPolicy="Equivalent",timeoutSeconds=10
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="connectorbindings",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE","DELETE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type ConnectorBindingWebhook struct {
	client ctrlcli.Client
}

func (r *ConnectorBindingWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	r.client = opts.Manager.GetClient()

	return &walruscore.ConnectorBinding{}, nil
}

var _ ctrlwebhook.CustomValidator = (*ConnectorBindingWebhook)(nil)

func (r *ConnectorBindingWebhook) ValidateCreate(
	ctx context.Context,
	obj runtime.Object,
) (ctrladmission.Warnings, error) {
	cb := obj.(*walruscore.ConnectorBinding)

	conn := &walruscore.Connector{
		ObjectMeta: meta.ObjectMeta{
			Name:      cb.Spec.Connector.Name,
			Namespace: cb.Spec.Connector.Namespace,
		},
	}
	err := r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(conn), conn)
	if err != nil {
		return nil, err
	}

	// Validate connector applicable environment type.
	{
		envList := new(walrus.EnvironmentList)
		err = r.client.List(ctx, envList, ctrlcli.MatchingFields{"metadata.name": cb.Namespace})
		if err != nil {
			return nil, err
		}

		if len(envList.Items) != 1 {
			return nil, field.Forbidden(
				field.NewPath("metadata", "namespace"), "connector must be created to a valid environment")
		}

		if envList.Items[0].Spec.Type != conn.Spec.ApplicableEnvironmentType {
			return nil, field.Invalid(
				field.NewPath("spec", "connector"),
				cb.Spec.Connector,
				fmt.Sprintf("connector must be created to a %s environment", conn.Spec.ApplicableEnvironmentType),
			)
		}
	}

	// Validate duplicated binding.
	{
		cbList := new(walruscore.ConnectorBindingList)
		err := r.client.List(ctx, cbList, ctrlcli.InNamespace(cb.Namespace))
		if err != nil {
			return nil, err
		}

		for i := range cbList.Items {
			if cbList.Items[i].Status.Category == conn.Spec.Category &&
				cbList.Items[i].Status.Type == conn.Spec.Type {
				return nil, field.Invalid(
					field.NewPath("spec", "connector"),
					cb.Spec.Connector,
					"connectors for the same purpose cannot be bound repeatedly",
				)
			}
		}
	}

	return nil, nil
}

func (r *ConnectorBindingWebhook) ValidateUpdate(
	ctx context.Context,
	oldObj, newObj runtime.Object,
) (ctrladmission.Warnings, error) {
	oldCb, newCb := oldObj.(*walruscore.ConnectorBinding), newObj.(*walruscore.ConnectorBinding)
	if !reflect.DeepEqual(oldCb.Spec, newCb.Spec) {
		return nil, field.Forbidden(field.NewPath("spec"), "cannot update connector binding spec")
	}

	return nil, nil
}

func (r *ConnectorBindingWebhook) ValidateDelete(
	ctx context.Context,
	obj runtime.Object,
) (ctrladmission.Warnings, error) {
	return nil, nil
}

func (r *ConnectorBindingWebhook) Default(ctx context.Context, obj runtime.Object) error {
	cb := obj.(*walruscore.ConnectorBinding)
	if cb.DeletionTimestamp != nil {
		return nil
	}

	connector := &walruscore.Connector{
		ObjectMeta: meta.ObjectMeta{
			Name:      cb.Spec.Connector.Name,
			Namespace: cb.Spec.Connector.Namespace,
		},
	}
	if err := r.client.Get(ctx, ctrlcli.ObjectKeyFromObject(connector), connector); err != nil {
		return err
	}

	labels := cb.Labels
	if labels == nil {
		labels = map[string]string{}
	}
	labels["walrus.seal.io/connector"] = fmt.Sprintf("%s-%s", cb.Spec.Connector.Namespace, cb.Spec.Connector.Name)
	cb.SetLabels(labels)

	return nil
}
