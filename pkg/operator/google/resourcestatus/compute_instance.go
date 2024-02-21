package resourcestatus

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/api/compute/v1"

	"github.com/seal-io/walrus/pkg/dao/types/status"
	gtypes "github.com/seal-io/walrus/pkg/operator/google/types"
	"github.com/seal-io/walrus/pkg/operator/types"
)

func getComputeInstance(ctx context.Context, resourceType, name string) (*status.Status, error) {
	cred, ok := ctx.Value(types.CredentialKey).(*gtypes.Credential)
	if !ok {
		return nil, errors.New("not found credential from context")
	}

	service, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}

	instance, err := service.Instances.Get(cred.Project, cred.Zone, name).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get google resource %s %s: %w", resourceType, name, err)
	}

	return computeInstanceStatusConverter.Convert(strings.ToLower(instance.Status), instance.StatusMessage), nil
}
