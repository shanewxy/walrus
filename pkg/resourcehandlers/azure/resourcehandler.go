package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/seal-io/utils/stringx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	aztypes "github.com/seal-io/walrus/pkg/resourcehandlers/azure/types"
	"github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

const OperatorType = resourcehandler.ConnectorTypeAzure

// New returns resourcehandlers.ResourceHandler with the given options.
func New(ctx context.Context, opts resourcehandler.CreateOptions) (resourcehandler.ResourceHandler, error) {
	name := opts.Connector.Name
	config, err := types.GetConfigData(ctx, opts)
	if err != nil {
		return nil, err
	}

	cred, err := aztypes.GetCredential(config)
	if err != nil {
		return nil, err
	}

	return Operator{
		name:       name,
		cred:       cred,
		identifier: stringx.SumBySHA256("azure:", cred.SubscriptionID, cred.TenantID, cred.ClientID),
	}, nil
}

type Operator struct {
	name       string
	cred       *aztypes.Credential
	identifier string
}

func (o Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (o Operator) IsConnected(ctx context.Context) error {
	cred, err := azidentity.NewClientSecretCredential(o.cred.TenantID, o.cred.ClientID, o.cred.ClientSecret, nil)
	if err != nil {
		return err
	}

	clientFactory, err := armresources.NewClientFactory(o.cred.SubscriptionID, cred, nil)
	if err != nil {
		return err
	}

	client := clientFactory.NewResourceGroupsClient()

	pager := client.NewListPager(nil)

	_, err = pager.NextPage(ctx)
	if err != nil {
		return fmt.Errorf("error connect to azure: %w", err)
	}

	return nil
}

func (o Operator) Burst() int {
	return 100
}

// ID implements resourcehandlers.ResourceHandler.
func (o Operator) ID() string {
	return o.identifier
}

// GetComponents implements resourcehandlers.ResourceHandler.
func (o Operator) GetComponents(
	ctx context.Context,
	resource *walrus.ResourceComponents,
) ([]*walrus.ResourceComponents, error) {
	return nil, nil
}

// Log implements resourcehandlers.ResourceHandler.
func (o Operator) Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error {
	return nil
}

// Exec implements resourcehandlers.ResourceHandler.
func (o Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	return nil
}

func (o Operator) GetKeys(ctx context.Context, component *walrus.ResourceComponents) (*resourcehandler.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (o Operator) GetStatus(ctx context.Context, component *walrus.ResourceComponents) ([]meta.Condition, error) {
	// TODO: Implement this method after resource is migrated.

	return nil, nil
}
