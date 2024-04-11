package apistatus

import "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	CatalogConditionReady   v1.ConditionType = "Ready"
	CatalogConditionRefresh v1.ConditionType = "Refresh"
)

const (
	CatalogConditionReasonPreparing = "Preparing"
	CatalogConditionReasonError     = "Error"
)

// catalogStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Refresh          | Unknown                 | Refreshing            | Transitioning         |
//	| Refresh          | False                   | /                     | /                     |
//	| Refresh          | True                    | /                     | /                     |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Interrupted           |
//	| Ready            | True                    | Ready                 |                       |
var catalogStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			CatalogConditionRefresh,
			CatalogConditionReady,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(CatalogConditionRefresh,
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

func WalkCatalog(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return catalogStatusPaths.Walk(st)
}
