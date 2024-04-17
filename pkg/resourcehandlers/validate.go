package resourcehandlers

import (
	"context"
	"errors"
	"fmt"

	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
)

func IsConnected(ctx context.Context, conn *walruscore.Connector, client ctrlcli.Client) error {
	switch conn.Spec.Category {
	case walruscore.ConnectorCategoryKubernetes, walruscore.ConnectorCategoryCloudProvider:
		op, err := Get(ctx, resourcehandler.CreateOptions{
			Connector: *conn,
		})
		if err != nil {
			return err
		}

		if err = op.IsConnected(ctx); err != nil {
			return fmt.Errorf("unreachable connector: %w", err)
		}

	case walruscore.ConnectorCategoryCustom:

	default:
		return errors.New("invalid connector category")
	}

	return nil
}
