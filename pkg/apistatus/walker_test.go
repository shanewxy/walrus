package apistatus

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

func TestWalker_sxs(t *testing.T) {
	// 1. Define resource with status.
	type ExampleResource struct {
		Status walruscore.StatusDescriptor
	}

	// 2. Define the condition types of ExampleResource,
	// condition type can be past tense or present tense.
	const (
		ExampleResourceStatusProgressing    walruscore.ConditionType = "Progressing"
		ExampleResourceStatusReplicaFailure walruscore.ConditionType = "ReplicaFailure"
		ExampleResourceStatusAvailable      walruscore.ConditionType = "Available"
	)

	// 2.1  clarify the condition type and its status meaning as below.
	//      | Condition Type |     Condition Status    | Human Readable Status | Human Sensible Status |
	//      | -------------- | ----------------------- | --------------------- | --------------------- |
	//      | Progressing    | Unknown                 | Progressing           | PhaseIsTransitioning         |
	//      | Progressing    | False                   | Progressing           | Error                 |
	//      | Progressing    | True(ReplicaSetUpdated) | Progressing           | PhaseIsTransitioning         |
	//      | Progressing    | True(DeploymentPaused)  | Pausing               | PhaseIsTransitioning         |
	//      | Progressing    | True                    | Progressed            | Done                  |
	//      | ReplicaFailure | Unknown                 | ReplicaDeploying      | PhaseIsTransitioning         |
	//      | ReplicaFailure | False                   | ReplicaDeployed       | Done                  |
	//      | ReplicaFailure | True                    | ReplicaDeployFailed   | Error                 |
	//      | Available      | Unknown                 | Preparing             | PhaseIsTransitioning         |
	//      | Available      | False                   | Unavailable           | Error                 |
	//      | Available      | True                    | Available             | Done                  |

	// 3. Create a flow to connect the above condition types.
	f := NewWalker(
		// Define paths.
		[][]walruscore.ConditionType{
			{
				ExampleResourceStatusProgressing,
				ExampleResourceStatusReplicaFailure,
				ExampleResourceStatusAvailable,
			},
		},
		// Arrange the default step decision logic.
		func(d Decision[walruscore.ConditionType]) {
			d.Make(ExampleResourceStatusProgressing,
				func(st walruscore.ConditionStatus, reason string) (string, string, Score) {
					if st == walruscore.ConditionTrue && reason != "ReplicaSetUpdated" {
						return "Progressed", "", SummaryScoreDone
					}

					if st == walruscore.ConditionUnknown && reason == "DeploymentPaused" {
						return "Pausing", "", SummaryScoreTransitioning
					}

					return "Progressing", "", SummaryScoreTransitioning
				})

			d.Make(ExampleResourceStatusReplicaFailure,
				func(st walruscore.ConditionStatus, reason string) (string, string, Score) {
					switch st {
					case walruscore.ConditionFalse:
						return "ReplicaDeployed", "", SummaryScoreDone
					case walruscore.ConditionTrue:
						return "ReplicaDeployed", "", SummaryScoreInterrupted
					}

					return "ReplicaDeploying", "", SummaryScoreTransitioning
				})
		},
	)

	var p printer

	// 4. Create an instance of ExampleResource.
	var r ExampleResource
	// 4.1  at beginning, the status is empty(we haven't configured any conditions or summary result),
	//      the path will walk to the end step and display the info of the last step,
	//      so we should get a done available summary,
	//      which can treat as Default Status.
	p.Dump("Default Available [D]", f.Walk(&r.Status))
	// 4.2  marked the "Progressing" status to Unknown, which means progressing,
	//      we should get a transitioning progressing summary.
	ExampleResourceStatusProgressing.Unknown(&r, "", "")
	p.Dump("Progressing [T]", f.Walk(&r.Status))
	// 4.3  marked the "Progressing" status to True with ReplicaSetUpdated reason,
	//      we should still get a transitioning progressing summary.
	r.Status.Conditions[0].Status = walruscore.ConditionTrue
	r.Status.Conditions[0].Reason = "ReplicaSetUpdated"
	p.Dump("Still Progressing [T]", f.Walk(&r.Status))
	// 4.4  marked the "Progressing" reason to NewReplicaSetAvailable,
	//      we should get a done progressing summary.
	//      at the same time, we haven't configured other conditions,
	//      so we only can see the progressing result.
	r.Status.Conditions[0].Reason = "NewReplicaSetAvailable"
	p.Dump("Progressed [D]", f.Walk(&r.Status))
	// 4.5  marked the "ReplicaFailure" status to Unknown, which means replica deploying,
	//      we should get a transitioning replica deploying summary.
	ExampleResourceStatusReplicaFailure.Unknown(&r, "", "")
	p.Dump("Replica Deploying [T]", f.Walk(&r.Status))
	// 4.6  marked the "ReplicaFailure" status to True, which means replica deploying failed,
	//      we should get a failed replica deploy summary.
	ExampleResourceStatusReplicaFailure.True(&r, "", "")
	p.Dump("Replica Deploy Failed [E]", f.Walk(&r.Status))
	// 4.7  marked the "Available" status to Unknown,
	//      we still get a failed replica deploy summary,
	//      as the path cannot move the next step as the "ReplicaFailure" step is not False.
	ExampleResourceStatusAvailable.Unknown(&r, "", "")
	p.Dump("Still Replica Deploy Failed [E]", f.Walk(&r.Status))
	// 4.8  until marked the "ReplicaFailure" status to False or remove "ReplicaFailure" condition,
	//      we will get a transitioning preparing summary.
	ExampleResourceStatusReplicaFailure.False(&r, "", "")
	p.Dump("Preparing [T]", f.Walk(&r.Status))
	// 4.9  marked the "Available" status to False, which means replica deploying failed,
	//      we should get an error unavailable summary.
	ExampleResourceStatusAvailable.False(&r, "", "")
	p.Dump("Unavailable [E]", f.Walk(&r.Status))
	// 4.10 marked the "Progressing" status to Unknown, which means progressing again,
	//      we should get a transitioning progressing summary.
	ExampleResourceStatusProgressing.Unknown(&r, "", "")
	p.Dump("Progressing Again [T]", f.Walk(&r.Status))

	t.Log(p.String())
}

func TestWalker_multiple(t *testing.T) {
	const (
		ExampleResourceStatusDeployed walruscore.ConditionType = "Deployed"
		ExampleResourceStatusReady    walruscore.ConditionType = "Ready"
		ExampleResourceStatusDeleted  walruscore.ConditionType = "Deleted"
	)

	f := NewWalker(
		[][]walruscore.ConditionType{
			{
				ExampleResourceStatusDeployed,
				ExampleResourceStatusReady,
			},
			{
				ExampleResourceStatusDeleted,
			},
		},
		func(d Decision[walruscore.ConditionType]) {
			d.Make(ExampleResourceStatusDeleted,
				func(st walruscore.ConditionStatus, reason string) (string, string, Score) {
					switch st {
					case walruscore.ConditionUnknown:
						return "Deleting", "", SummaryScoreTransitioning
					case walruscore.ConditionFalse:
						return "DeleteFailed", "", SummaryScoreInterrupted
					}

					return "Deleted", "", SummaryScoreDone
				})
		},
	)

	type (
		input struct {
			Status walruscore.StatusDescriptor
			Before func(*input)
		}
	)

	testCases := []struct {
		name     string
		given    input
		expected *walruscore.ConditionSummary
	}{
		{
			name: "no conditions",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: nil,
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "Ready",
			},
		},
		{
			name: "first deploy",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionUnknown,
						},
					},
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "Deploying",
			},
		},
		{
			name: "deployed",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusReady,
							Status: walruscore.ConditionUnknown,
						},
					},
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "Preparing",
			},
		},
		{
			name: "redeploy",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionUnknown,
						},
						{
							Type:   ExampleResourceStatusReady,
							Status: walruscore.ConditionTrue,
						},
					},
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "Deploying",
			},
		},
		{
			name: "redeploy but failed",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionFalse,
						},
						{
							Type:   ExampleResourceStatusReady,
							Status: walruscore.ConditionTrue,
						},
					},
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "DeployFailed",
			},
		},
		{
			name: "delete",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusReady,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusDeleted,
							Status: walruscore.ConditionUnknown,
						},
					},
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "Deleting",
			},
		},
		{
			name: "delete but failed",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusReady,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusDeleted,
							Status: walruscore.ConditionFalse,
						},
					},
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "DeleteFailed",
			},
		},
		{
			name: "delete failed but redeploy",
			given: input{
				Status: walruscore.StatusDescriptor{
					Conditions: []walruscore.Condition{
						{
							Type:   ExampleResourceStatusDeployed,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusReady,
							Status: walruscore.ConditionTrue,
						},
						{
							Type:   ExampleResourceStatusDeleted,
							Status: walruscore.ConditionFalse,
						},
					},
				},
				Before: func(i *input) {
					// Remove deleted status and mark deployed status.
					ExampleResourceStatusDeployed.Reset(i, "")
				},
			},
			expected: &walruscore.ConditionSummary{
				Phase: "Deploying",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.given.Before != nil {
				tc.given.Before(&tc.given)
			}
			actual := f.Walk(&tc.given.Status)
			assert.Equal(t, tc.expected, actual, "case %q", tc.name)
		})
	}
}

type printer struct {
	sb strings.Builder
}

func (p *printer) Dump(title string, s *walruscore.ConditionSummary) {
	p.sb.WriteString(title)
	p.sb.WriteString(": ")
	spew.Fdump(&p.sb, s)
	p.sb.WriteString("\n")
}

func (p *printer) String() string {
	return p.sb.String()
}
