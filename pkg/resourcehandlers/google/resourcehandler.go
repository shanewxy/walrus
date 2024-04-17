package google

import (
	"context"
	"fmt"

	"github.com/seal-io/utils/stringx"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	gtypes "github.com/seal-io/walrus/pkg/resourcehandlers/google/types"
	"github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

const OperatorType = resourcehandler.ConnectorTypeGoogle

// New returns resourcehandlers.ResourceHandler with the given options.
func New(ctx context.Context, opts resourcehandler.CreateOptions) (resourcehandler.ResourceHandler, error) {
	name := opts.Connector.Name
	config, err := types.GetConfigData(ctx, opts)
	if err != nil {
		return nil, err
	}

	cred, err := gtypes.GetCredential(config)
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       name,
		cred:       cred,
		identifier: stringx.SumBySHA256("google:", cred.Project, cred.Region, cred.Zone),
	}, nil
}

type Operator struct {
	name       string
	cred       *gtypes.Credential
	identifier string
}

func (op Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (op Operator) IsConnected(ctx context.Context) error {
	service, err := compute.NewService(ctx, option.WithCredentialsJSON([]byte(op.cred.Credentials)))
	if err != nil {
		return err
	}

	_, err = service.Regions.List(op.cred.Project).Do()
	if err != nil {
		return fmt.Errorf("error connect to google cloud: %w", err)
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
