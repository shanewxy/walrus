package types

import (
	"context"

	"github.com/seal-io/walrus/pkg/resourcehandler"
)

type ExecutableResource interface {
	Exec(ctx context.Context, key string, opts resourcehandler.ExecOptions) error
	Supported(ctx context.Context, key string) (bool, error)
}

type LoggableResource interface {
	Log(ctx context.Context, key string, opts resourcehandler.LogOptions) error
}
