package k8s

import (
	"context"
	"time"

	"github.com/seal-io/utils/stringx"
	"github.com/seal-io/utils/waitx"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	dynamicclient "k8s.io/client-go/dynamic"
	batchclient "k8s.io/client-go/kubernetes/typed/batch/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	networkingclient "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
)

const OperatorType = resourcehandler.ConnectorTypeKubernetes

// New returns resourcehandlers.ResourceHandler with the given options.
func New(ctx context.Context, opts resourcehandler.CreateOptions) (resourcehandler.ResourceHandler, error) {
	restConfig, err := GetConfig(ctx, opts, func(c *rest.Config) {
		c.Timeout = 0
	})
	if err != nil {
		return nil, err
	}

	restCli, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, err
	}

	coreCli, err := coreclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	batchCli, err := batchclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	networkingCli, err := networkingclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	dynamicCli, err := dynamicclient.NewForConfigAndClient(restConfig, restCli)
	if err != nil {
		return nil, err
	}

	op := Operator{
		Logger:        klog.Background().WithName("resourcehandlers").WithName("k8s"),
		Identifier:    stringx.SumBySHA256("k8s:", restConfig.Host, restConfig.APIPath),
		RestConfig:    restConfig,
		CoreCli:       coreCli,
		BatchCli:      batchCli,
		NetworkingCli: networkingCli,
		DynamicCli:    dynamicCli,
	}

	return op, nil
}

type Operator struct {
	Logger        klog.Logger
	Identifier    string
	RestConfig    *rest.Config
	CoreCli       *coreclient.CoreV1Client
	BatchCli      *batchclient.BatchV1Client
	NetworkingCli *networkingclient.NetworkingV1Client
	DynamicCli    *dynamicclient.DynamicClient
}

// Type implements resourcehandlers.ResourceHandler.
func (Operator) Type() resourcehandler.Type {
	return OperatorType
}

// IsConnected implements resourcehandlers.ResourceHandler.
func (op Operator) IsConnected(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := waitx.PollUntilContextCancel(ctx, time.Second, true,
		func(ctx context.Context) error {
			return IsConnected(context.TODO(), op.CoreCli.RESTClient())
		},
	)

	return err
}

// Burst implements resourcehandlers.ResourceHandler.
func (op Operator) Burst() int {
	if op.RestConfig.Burst == 0 {
		return rest.DefaultBurst
	}

	return op.RestConfig.Burst
}

// ID implements resourcehandlers.ResourceHandler.
func (op Operator) ID() string {
	return op.Identifier
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
