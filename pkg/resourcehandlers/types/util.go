package types

import (
	"context"
	"fmt"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/system"
)

func GetConfigData(ctx context.Context, opts resourcehandler.CreateOptions) (map[string][]byte, error) {
	con := opts.Connector

	sec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: con.Namespace,
			Name:      con.Spec.SecretName,
		},
	}

	cli := system.LoopbackCtrlClient.Get()
	err := cli.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec)
	if err != nil {
		return nil, fmt.Errorf("error get secret %s: %w", sec.Name, err)
	}
	return sec.Data, nil
}
