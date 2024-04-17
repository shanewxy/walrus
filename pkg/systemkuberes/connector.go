package systemkuberes

import (
	"context"
	"fmt"

	dtypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubeconfig"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/system"
)

func installDefaultKubernetesConnector(
	ctx context.Context,
	cli clientset.Interface,
	project string,
	envType walruscore.EnvironmentType,
) (*walrus.Connector, error) {
	connCli := cli.WalrusV1().Connectors(project)

	config, err := readKubeConfig()
	if err != nil {
		return nil, fmt.Errorf("read kube config: %w", err)
	}

	c := &walrus.Connector{
		ObjectMeta: meta.ObjectMeta{
			Namespace: project,
			Name:      DefaultConnectorName,
		},
		Spec: walruscore.ConnectorSpec{
			Type:                      resourcehandler.ConnectorTypeKubernetes,
			Description:               "Local Kubernetes",
			Category:                  walruscore.ConnectorCategoryKubernetes,
			ApplicableEnvironmentType: envType,
			Config: walruscore.ConnectorConfig{
				Version: "v1",
				Data: map[string]walruscore.ConnectorConfigEntry{
					"kubeconfig": {
						Value:   config,
						Visible: false,
					},
				},
			},
		},
	}

	conn, err := kubeclientset.Update(ctx, connCli, c, kubeclientset.WithCreateIfNotExisted[*walrus.Connector]())
	if err != nil {
		return nil, fmt.Errorf("install default kubernetes connector: %w", err)
	}

	return conn, nil
}

func readKubeConfig() (string, error) {
	kubeConfig := system.LoopbackKubeClientConfig.Get()

	kc, err := kubeconfig.ConvertRestConfigToApiConfig(&kubeConfig)
	if err != nil {
		return "", err
	}

	kcData, err := clientcmd.Write(kc)
	if err != nil {
		return "", err
	}

	return string(kcData), err
}

func installDefaultDockerConnector(
	ctx context.Context,
	cli clientset.Interface,
	project string,
	envType walruscore.EnvironmentType,
) (*walrus.Connector, error) {
	connCli := cli.WalrusV1().Connectors(project)

	c := &walrus.Connector{
		ObjectMeta: meta.ObjectMeta{
			Namespace: project,
			Name:      DefaultConnectorName,
		},
		Spec: walruscore.ConnectorSpec{
			Type:                      resourcehandler.ConnectorTypeDocker,
			Category:                  walruscore.ConnectorCategoryDocker,
			ApplicableEnvironmentType: envType,
			Config: walruscore.ConnectorConfig{
				Version: "v1",
				Data:    map[string]walruscore.ConnectorConfigEntry{},
			},
		},
	}

	conn, err := kubeclientset.Create(ctx, connCli, c)
	if err != nil {
		return nil, fmt.Errorf("install default docker connector: %w", err)
	}

	if err := applyLocalDockerNetwork(ctx); err != nil {
		return nil, fmt.Errorf("apply local docker network: %w", err)
	}

	return conn, nil
}

func applyLocalDockerNetwork(ctx context.Context) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	networkName := "local-walrus"

	networks, err := cli.NetworkList(ctx, dtypes.NetworkListOptions{})
	if err != nil {
		return err
	}

	exists := false

	for _, n := range networks {
		if n.Name == networkName {
			exists = true
			break
		}
	}

	if !exists {
		_, err = cli.NetworkCreate(ctx, networkName, dtypes.NetworkCreate{
			Driver: "bridge",
		})
		if err != nil {
			return err
		}
	}

	return nil
}
