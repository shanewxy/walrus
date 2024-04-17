package alibaba

import (
	"context"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba/resourceexec"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba/resourcelog"
	"github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

const OperatorType = resourcehandler.ConnectorTypeAlibabaCloud

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
		identifier: stringx.SumBySHA256("alibaba:", cred.AccessKey, cred.AccessSecret),
	}, nil
}

type Operator struct {
	name       string
	cred       *types.Credential
	identifier string
}

func (op Operator) IsConnected(ctx context.Context) error {
	client, err := ecs.NewClientWithAccessKey(
		op.cred.Region,
		op.cred.AccessKey,
		op.cred.AccessSecret,
	)
	if err != nil {
		return fmt.Errorf("error create alibaba client %s: %w", op.name, err)
	}

	// Use DescribeRegion API to check reachable and user has access to region.
	// https://www.alibabacloud.com/help/en/elastic-compute-service/latest/regions-describeregions
	req := ecs.CreateDescribeRegionsRequest()
	req.Scheme = "HTTPS"

	_, err = client.DescribeRegions(req)
	if err != nil {
		return fmt.Errorf("error connect to %s: %w", op.name, err)
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
