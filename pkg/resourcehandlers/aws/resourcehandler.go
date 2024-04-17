package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/aws/resourceexec"
	"github.com/seal-io/walrus/pkg/resourcehandlers/aws/resourcelog"
	opawstypes "github.com/seal-io/walrus/pkg/resourcehandlers/aws/types"
	"github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

const OperatorType = resourcehandler.ConnectorTypeAWS

// New returns resourcehandlers.ResourceHandler with the given options.
func New(ctx context.Context, opts resourcehandler.CreateOptions) (resourcehandler.ResourceHandler, error) {
	name := opts.Connector.Name

	config, err := types.GetConfigData(ctx, opts)
	if err != nil {
		return nil, err
	}

	cred, err := types.GetCredential(config)
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       name,
		cred:       cred,
		identifier: stringx.SumBySHA256("aws:", cred.AccessKey, cred.AccessSecret),
	}, nil
}

type Operator struct {
	name       string
	cred       *types.Credential
	identifier string
}

func (op Operator) IsConnected(ctx context.Context) error {
	cred := opawstypes.Credential(*op.cred)

	cfg, err := cred.Config()
	if err != nil {
		return err
	}

	// Use DescribeRegions API to check reachable.
	cli := ec2.NewFromConfig(*cfg)

	_, err = cli.DescribeRegions(ctx, nil)
	if err != nil {
		return fmt.Errorf("error connect to aws: %w", err)
	}

	return nil
}

func (op Operator) Type() resourcehandler.Type {
	return OperatorType
}

// Burst implements resourcehandlers.ResourceHandler.
func (op Operator) Burst() int {
	return 200
}

// ID implements resourcehandlers.ResourceHandler.
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
	newCtx := context.WithValue(ctx, types.CredentialKey, op.cred)
	return resourcelog.Log(newCtx, key, opts)
}

// Exec implements resourcehandlers.ResourceHandler.
func (op Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	newCtx := context.WithValue(ctx, types.CredentialKey, op.cred)
	return resourceexec.Exec(newCtx, key, opts)
}

func (op Operator) GetKeys(ctx context.Context, component *walrus.ResourceComponents) (*resourcehandler.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (op Operator) GetStatus(ctx context.Context, component *walrus.ResourceComponents) ([]meta.Condition, error) {
	// TODO: Implement this method after resource is migrated.

	return nil, nil
}
