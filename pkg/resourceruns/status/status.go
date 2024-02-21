package status

import (
	"context"
	"fmt"

	runbus "github.com/seal-io/walrus/pkg/bus/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

func IsStatusRunning(run *model.ResourceRun) bool {
	return status.ResourceRunStatusPending.IsUnknown(run) ||
		status.ResourceRunStatusPlanned.IsUnknown(run) ||
		status.ResourceRunStatusApply.IsUnknown(run)
}

func IsStatusFailed(run *model.ResourceRun) bool {
	return status.ResourceRunStatusPending.IsFalse(run) ||
		status.ResourceRunStatusPlanned.IsFalse(run) ||
		status.ResourceRunStatusApply.IsFalse(run)
}

func IsStatusPending(run *model.ResourceRun) bool {
	return status.ResourceRunStatusPending.IsUnknown(run)
}

func IsStatusPlanned(run *model.ResourceRun) bool {
	return status.ResourceRunStatusPlanned.IsTrue(run) &&
		!status.ResourceRunStatusApply.Exist(run)
}

// IsStatusPlanCondition checks if the resource run is in the plan condition.
func IsStatusPlanCondition(run *model.ResourceRun) bool {
	return status.ResourceRunStatusPlanned.Exist(run) &&
		!status.ResourceRunStatusApply.Exist(run)
}

func IsStatusSucceeded(run *model.ResourceRun) bool {
	return status.ResourceRunStatusApply.IsTrue(run)
}

// SetStatusFalse sets the status of the resource run to false.
func SetStatusFalse(run *model.ResourceRun, errMsg string) {
	switch {
	case status.ResourceRunStatusPlanned.IsUnknown(run):
		errMsg = fmt.Sprintf("plan failed: %s", errMsg)
		status.ResourceRunStatusPlanned.False(run, errMsg)
	case status.ResourceRunStatusApply.IsUnknown(run):
		errMsg = fmt.Sprintf("apply failed: %s", errMsg)
		status.ResourceRunStatusApply.False(run, errMsg)
	}

	run.Status.SetSummary(status.WalkResourceRun(&run.Status))
}

// SetStatusTrue sets the status of the resource run to true.
// It marks the status of the resource run as "Succeeded".
func SetStatusTrue(run *model.ResourceRun, msg string) {
	switch {
	case status.ResourceRunStatusPlanned.IsUnknown(run):
		status.ResourceRunStatusPlanned.True(run, msg)
	case status.ResourceRunStatusApply.IsUnknown(run):
		status.ResourceRunStatusApply.True(run, msg)
	}

	run.Status.SetSummary(status.WalkResourceRun(&run.Status))
}

// UpdateStatus updates the status of the resource run.
func UpdateStatus(ctx context.Context, mc model.ClientSet, run *model.ResourceRun) (*model.ResourceRun, error) {
	if run == nil {
		return nil, nil
	}

	// Report to resource run.
	run.Status.SetSummary(status.WalkResourceRun(&run.Status))

	run, err := mc.ResourceRuns().UpdateOne(run).
		SetStatus(run.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = runbus.Notify(ctx, mc, run)
	if err != nil {
		return nil, err
	}

	return run, nil
}