package leaderelection

import (
	"context"

	"github.com/seal-io/utils/json"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

func Dummy() resourcelock.Interface {
	return dummy{}
}

type dummy struct{}

func (d dummy) Get(ctx context.Context) (*resourcelock.LeaderElectionRecord, []byte, error) {
	r := resourcelock.LeaderElectionRecord{
		HolderIdentity:       "dummy-holder",
		LeaseDurationSeconds: 3600,
		AcquireTime:          meta.Now(),
		RenewTime:            meta.Now(),
		LeaderTransitions:    0,
	}
	rbs, err := json.Marshal(r)
	if err != nil {
		return nil, nil, err
	}
	return &r, rbs, nil
}

func (d dummy) Create(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	return nil
}

func (d dummy) Update(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	return nil
}

func (d dummy) RecordEvent(s string) {
}

func (d dummy) Identity() string {
	return "dummy"
}

func (d dummy) Describe() string {
	return "dummy/dummy"
}
