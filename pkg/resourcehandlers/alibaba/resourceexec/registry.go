package resourceexec

import (
	"context"
	"errors"

	"k8s.io/klog/v2"

	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba/key"
	"github.com/seal-io/walrus/pkg/resourcehandlers/types"
)

// resourceTypes indicate supported resource type and their functions.
var resourceTypes map[string]getExecutableResource

type getExecutableResource func(ctx context.Context) (types.ExecutableResource, error)

func init() {
	resourceTypes = map[string]getExecutableResource{
		"alicloud_instance": getEcsInstance,
	}
}

// Supported indicate whether the resource is supported to exec.
func Supported(ctx context.Context, k string) (bool, error) {
	resourceType, name, ok := key.Decode(k)
	if !ok {
		return false, errors.New("invalid key")
	}

	fs, exist := resourceTypes[resourceType]
	if !exist {
		return false, nil
	}

	res, err := fs(ctx)
	if err != nil {
		return false, err
	}

	supported, err := res.Supported(ctx, name)
	if err != nil {
		return false, err
	}

	if !supported {
		return false, nil
	}

	return supported, nil
}

// Exec resource by key.
func Exec(ctx context.Context, k string, opts resourcehandler.ExecOptions) error {
	supported, err := Supported(ctx, k)
	if err != nil {
		return err
	}

	if !supported {
		return errors.New("unsupported resource type")
	}

	resourceType, name, ok := key.Decode(k)
	if !ok {
		return errors.New("invalid key")
	}

	fs := resourceTypes[resourceType]

	res, err := fs(ctx)
	if err != nil {
		return err
	}

	err = res.Exec(ctx, name, opts)
	if err != nil {
		klog.Warningf("error exec resource %s/%s: %v", resourceType, name, err)
	}

	return err
}
