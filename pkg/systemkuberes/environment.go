package systemkuberes

import (
	"context"
	"fmt"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubeclientset/review"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

const (
	// DefaultEnvironmentName is the Kubernetes Namespace name for the default environment.
	DefaultEnvironmentName = DefaultProjectName + "-local"

	// DefaultConnectorName is the name of the default connector.
	DefaultConnectorName = "local"
)

// InstallDefaultEnvironment creates the default environment, alias to Kubernetes Namespace default-local.
func InstallDefaultEnvironment(ctx context.Context, cli clientset.Interface) error {
	localEnvironmentMode, err := systemsetting.DefaultEnvironmentMode.Value(ctx)
	if err != nil {
		return err
	}

	if localEnvironmentMode == "disabled" {
		return nil
	}

	err = review.CanDoCreate(ctx,
		cli.AuthorizationV1().SelfSubjectAccessReviews(),
		review.Simples{
			{
				Group:    walrus.SchemeGroupVersion.Group,
				Version:  walrus.SchemeGroupVersion.Version,
				Resource: "environments",
			},
		},
	)
	if err != nil {
		return err
	}

	environmentType := func() walruscore.EnvironmentType {
		if system.LoopbackKubeInside.Get() {
			return walruscore.EnvironmentTypeProduction
		}
		return walruscore.EnvironmentTypeDevelopment
	}()

	envCli := cli.WalrusV1().Environments(DefaultProjectName)
	env := &walrus.Environment{
		ObjectMeta: meta.ObjectMeta{
			Namespace: DefaultProjectName,
			Name:      DefaultEnvironmentName,
		},
		Spec: walrus.EnvironmentSpec{
			Type:        environmentType,
			DisplayName: "Default Environment",
			Description: "The default environment created by Walrus.",
		},
	}

	_, err = kubeclientset.Create(ctx, envCli, env)
	if err != nil {
		return fmt.Errorf("install default environment: %w", err)
	}

	// Install default connector.
	var conn *walrus.Connector
	switch localEnvironmentMode {
	case "kubernetes":
		conn, err = installDefaultKubernetesConnector(ctx, cli, DefaultProjectName, environmentType)
		if err != nil {
			return err
		}
	case "docker":
		conn, err = installDefaultDockerConnector(ctx, cli, DefaultProjectName, environmentType)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid local environment mode %q", localEnvironmentMode)
	}

	// Create connector binding.
	bindingsCli := cli.WalrusV1().ConnectorBindings(DefaultEnvironmentName)
	connectorBinding := &walrus.ConnectorBinding{
		ObjectMeta: meta.ObjectMeta{
			Namespace: DefaultEnvironmentName,
			Name:      DefaultConnectorName,
		},
		Spec: walruscore.ConnectorBindingSpec{
			Connector: walruscore.ConnectorReference{
				Name:      conn.Name,
				Namespace: conn.Namespace,
			},
		},
	}

	_, err = kubeclientset.Create(ctx, bindingsCli, connectorBinding)
	if err != nil {
		return fmt.Errorf("install default connector binding: %w", err)
	}

	return nil
}
