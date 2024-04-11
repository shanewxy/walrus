package apistatus

import "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	TemplateConditionReady   v1.ConditionType = "Ready"
	TemplateConditionRefresh v1.ConditionType = "Refresh"
)

const (
	TemplateConditionReasonPreparing = "Preparing"
	TemplateConditionReasonError     = "Error"
)

// templateStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Refresh          | Unknown                 | Refreshing            | Transitioning         |
//	| Refresh          | False                   | /                     | /                     |
//	| Refresh          | True                    | /                     | /                     |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Interrupted           |
//	| Ready            | True                    | Ready                 |                       |
var templateStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			TemplateConditionRefresh,
			TemplateConditionReady,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(TemplateConditionRefresh,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				default:
					return "Refresh", "", SummaryScoreDone
				case v1.ConditionUnknown:
					return "Refreshing", "", SummaryScoreTransitioning
				}
			})
	},
)

func WalkTemplate(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return templateStatusPaths.Walk(st)
}
