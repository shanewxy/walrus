package v1

// StatusDescriptor include conditions and it's summary.
type StatusDescriptor struct {
	ConditionSummary `json:",inline"`

	// Conditions holds the conditions for the object.
	Conditions []Condition `json:"conditions,omitempty"`
}

// ConditionSummary is the summary of conditions.
type ConditionSummary struct {
	// Phase is the summary of conditions.
	Phase string `json:"phase,omitempty"`

	// PhaseMessage is the message of the phase.
	PhaseMessage string `json:"phaseMessage,omitempty"`
}
