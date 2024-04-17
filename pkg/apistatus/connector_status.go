package apistatus

import v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	ConnectorConditionConnected v1.ConditionType = "Connected"
)

const (
	ConnectorConditionReasonConnecting   = "Connecting"
	ConnectorConditionReasonDisconnected = "Disconnected"
)

// connectorStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |\
//	| Connected        | Unknown                 | Connecting            | Transitioning         |
//	| Connected        | False                   | Disconnected          | Error                 |
//	| Connected        | True                    | Connected             |                       |
var connectorStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			ConnectorConditionConnected,
		},
	},
)

// WalkConnector walks the given status by connector flow.
func WalkConnector(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return connectorStatusPaths.Walk(st)
}
