package controllers

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/controllers/walrus"
	"github.com/seal-io/walrus/pkg/controllers/walruscore"
)

// NB(thxCode): Register controllers below.
var setupers = []controller.Setup{
	new(walrus.ProjectSubjectAuthzReconciler),
	new(walrus.SubjectAuthzReconciler),
	new(walruscore.CatalogReconciler),
	new(walruscore.ConnectorReconciler),
	new(walruscore.ConnectorBindingReconciler),
	new(walruscore.ResourceReconciler),
	new(walruscore.ResourceDefinitionReconciler),
	new(walruscore.TemplateReconciler),
}

// Setup installs the CRD controllers.
func Setup(ctx context.Context, mgr ctrl.Manager) error {
	for i := range setupers {
		opts := controller.SetupOptions{Manager: mgr}
		err := setupers[i].SetupController(ctx, opts)
		if err != nil {
			return fmt.Errorf("controller setup: %s: %w", spew.Sdump(setupers[i]), err)
		}
	}
	return nil
}
