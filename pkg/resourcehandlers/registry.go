package resourcehandlers

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/alibaba"
	"github.com/seal-io/walrus/pkg/resourcehandlers/aws"
	"github.com/seal-io/walrus/pkg/resourcehandlers/azure"
	"github.com/seal-io/walrus/pkg/resourcehandlers/docker"
	"github.com/seal-io/walrus/pkg/resourcehandlers/google"
	"github.com/seal-io/walrus/pkg/resourcehandlers/k8s"
	"github.com/seal-io/walrus/pkg/resourcehandlers/unknown"
)

var opOperators map[resourcehandler.Type]resourcehandler.Creator

func init() {
	opOperators = map[resourcehandler.Type]resourcehandler.Creator{
		// Register resourcehandlers creators as below.
		k8s.OperatorType:     k8s.New,
		aws.OperatorType:     aws.New,
		alibaba.OperatorType: alibaba.New,
		azure.OperatorType:   azure.New,
		google.OperatorType:  google.New,
		docker.OperatorType:  docker.New,
	}
}

// Get returns ResourceHandler with the given CreateOptions.
func Get(ctx context.Context, opts resourcehandler.CreateOptions) (op resourcehandler.ResourceHandler, err error) {
	f, exist := opOperators[opts.Connector.Spec.Type]
	if !exist {
		// Try to create an any resourcehandlers.
		op, err = unknown.New(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("unknown resourcehandlers: %s", opts.Connector.Spec.Type)
		}
	} else {
		op, err = f(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("error connecting %s resourcehandlers: %w", opts.Connector.Spec.Type, err)
		}
	}

	return op, nil
}
