package unknown

import (
	"context"
	"errors"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
)

const OperatorType = "Unknown"

// New returns types.ResourceHandler with the given options.
func New(ctx context.Context, opts resourcehandler.CreateOptions) (resourcehandler.ResourceHandler, error) {
	if opts.Connector.Spec.Category != walruscore.ConnectorCategoryCustom {
		return nil, errors.New("not custom connector")
	}

	return Operator{}, nil
}

type Operator struct{}

func (Operator) Type() resourcehandler.Type {
	return OperatorType
}

func (Operator) IsConnected(ctx context.Context) error {
	return nil
}

func (op Operator) Burst() int {
	return 100
}

func (op Operator) ID() string {
	return ""
}

func (op Operator) GetComponents(
	ctx context.Context,
	resource *walrus.ResourceComponents,
) ([]*walrus.ResourceComponents, error) {
	return nil, nil
}

func (Operator) Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error {
	return nil
}

func (Operator) Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error {
	return nil
}

func (op Operator) GetKeys(ctx context.Context, component *walrus.ResourceComponents) (*resourcehandler.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (op Operator) GetStatus(ctx context.Context, component *walrus.ResourceComponents) ([]meta.Condition, error) {
	// TODO: Implement this method after resource is migrated.

	return nil, nil
}
