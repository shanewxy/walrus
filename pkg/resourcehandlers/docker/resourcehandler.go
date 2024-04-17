package docker

import (
	"context"
	"errors"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	optypes "github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

const OperatorType = resourcehandler.ConnectorTypeDocker

// New returns resourcehandlers.ResourceHandler with the given options.
func New(ctx context.Context, opts resourcehandler.CreateOptions) (resourcehandler.ResourceHandler, error) {
	name := opts.Connector.Name
	config, err := optypes.GetConfigData(ctx, opts)
	if err != nil {
		return nil, err
	}

	host := string(config["host"])
	if host == "" {
		return nil, errors.New("host is empty")
	}

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       name,
		identifier: stringx.SumBySHA256("docker:", host),
		client:     cli,
	}, nil
}

type Operator struct {
	name       string
	identifier string
	client     *client.Client
}

func (op Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (op Operator) IsConnected(ctx context.Context) error {
	if _, err := op.client.ServerVersion(ctx); err != nil {
		return fmt.Errorf("error connect to docker daemon: %w", err)
	}

	return nil
}

func (op Operator) Burst() int {
	return 100
}

func (op Operator) ID() string {
	return op.identifier
}

// GetComponents implements resourcehandlers.ResourceHandler.
func (op Operator) GetComponents(
	ctx context.Context,
	resource *walrus.ResourceComponents,
) ([]*walrus.ResourceComponents, error) {
	return nil, nil
}

// Log implements resourcehandlers.ResourceHandler.
func (op Operator) Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error {
	return nil
}

// Exec implements resourcehandlers.ResourceHandler.
func (op Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	return nil
}

func (op Operator) GetKeys(ctx context.Context, component *walrus.ResourceComponents) (*resourcehandler.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (op Operator) GetStatus(ctx context.Context, component *walrus.ResourceComponents) ([]meta.Condition, error) {
	// TODO: Implement this method after resource is migrated.

	return nil, nil
}
