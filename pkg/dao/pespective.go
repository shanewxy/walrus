package dao

import (
	"errors"

	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/perspective"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

func PerspectiveCreates(mc model.ClientSet, input ...*model.Perspective) ([]*model.PerspectiveCreate, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid input: empty list")
	}

	var rrs = make([]*model.PerspectiveCreate, len(input))
	for i, r := range input {
		if r == nil {
			return nil, errors.New("invalid input: nil entity")
		}

		// required.
		var c = mc.Perspectives().Create().
			SetName(r.Name).
			SetStartTime(r.StartTime).
			SetEndTime(r.EndTime).
			SetBuiltin(r.Builtin)

		// optional.
		if len(r.AllocationQueries) != 0 {
			c.SetAllocationQueries(r.AllocationQueries)
		}

		rrs[i] = c
	}
	return rrs, nil
}

func PerspectiveUpdate(mc model.ClientSet, input *model.Perspective) (*model.PerspectiveUpdateOne, error) {
	if input == nil {
		return nil, errors.New("invalid input: nil entity")
	}
	// predicated.
	var ps = []predicate.Perspective{
		perspective.ID(input.ID),
	}
	var c = mc.Perspectives().UpdateOne(input).
		Where(ps...)
	if input.StartTime != "" {
		c.SetStartTime(input.StartTime)
	}
	if input.EndTime != "" {
		c.SetEndTime(input.EndTime)
	}
	if len(input.AllocationQueries) != 0 {
		c.SetAllocationQueries(input.AllocationQueries)
	}
	return c, nil
}