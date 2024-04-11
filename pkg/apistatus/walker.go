package apistatus

import (
	"strings"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

type (
	// Walker holds the steps and makes a summary of the status's conditions.
	Walker interface {
		// Walk walks all conditions of the status and gives a proper summary.
		Walk(*walruscore.StatusDescriptor) *walruscore.ConditionSummary
	}

	// Decide returns readable and sensible status by the given condition status and reason,
	// and moves to next path step if both returning `isError` and `isTransitioning` are false.
	Decide func(st walruscore.ConditionStatus, reason string) (phase, message string, score Score)
)

// NewWalker creates a stacking walker by the given steps group,
// and applies the customized decision to the steps group.
//   - `stepsGroup` specifies the path steps in line,
//     logically, move to next step if the current step is done.
//     By default, Walker decides to move to the next step on whether the corresponding condition is True status.
//   - `arrange` applies the customized decision,
//     for example, moving to next step on a dedicated step if its status is False,
//     or changing step's display content by its status.
func NewWalker[T ~string](stepsGroup [][]T, arranges ...func(Decision[T])) Walker {
	if len(stepsGroup) == 0 {
		panic("empty steps group")
	}

	fs := make(paths[T], 0, len(stepsGroup))
	for i := range stepsGroup {
		fs = append(fs, newPath(stepsGroup[i], arranges...))
	}

	return fs
}

// paths stacks a collection of path,
// and picks the highest score result.
type paths[T ~string] []path[T]

func (ps paths[T]) Walk(st *walruscore.StatusDescriptor) *walruscore.ConditionSummary {
	var r *ConditionSummary
	for i := range ps {
		l := ps[i].Walk(st)
		if r == nil {
			r = l
			continue
		}

		// Accept the result that has a higher score.
		ls, rs := l.Score, r.Score
		if ls <= rs {
			continue
		}
		r, rs = l, ls

		// Quit soon if found one highest result.
		if rs == HighestSummaryScore {
			break
		}
	}

	return &r.ConditionSummary
}

type Score int

const (
	SummaryScoreDone Score = iota
	SummaryScoreTransitioning
	SummaryScoreInterrupted
	HighestSummaryScore = SummaryScoreInterrupted
)

// ConditionSummary is the summary of conditions with score.
type ConditionSummary struct {
	walruscore.ConditionSummary
	Score Score
}

// newPath creates a path and initializes it.
func newPath[T ~string](steps []T, arranges ...func(Decision[T])) path[T] {
	if len(steps) == 0 {
		panic("empty steps")
	}

	p := path[T]{
		steps:       steps,
		stepsIndex:  make(map[T]int, len(steps)),
		stepsDecide: make([]Decide, len(steps)),
	}
	for i := range steps {
		// Loop check, panic if found.
		if _, exist := p.stepsIndex[steps[i]]; exist {
			panic("found loop")
		}
		p.stepsIndex[steps[i]] = i
		p.stepsDecide[i] = getGeneralDecide(steps[i])
	}

	// Change the default decide logic after arranging.
	for i := range arranges {
		arranges[i](Decision[T](p))
	}

	return p
}

// path holds the steps and makes a summary of the status's conditions.
type path[T ~string] struct {
	steps       []T
	stepsIndex  map[T]int
	stepsDecide []Decide
}

func (f path[T]) Walk(st *walruscore.StatusDescriptor) *ConditionSummary {
	s := &ConditionSummary{}

	// Walk the status if condition list is not empty.
	if len(st.Conditions) != 0 {
		// Map conditions with the specified steps for quick indexing.
		stepsConditionIndex := make([]int, len(f.steps))

		for i, c := range st.Conditions {
			// Plus 1 to avoid aligning not found item.
			if idx, exist := f.stepsIndex[T(c.Type)]; exist {
				stepsConditionIndex[idx] = i + 1
			}
		}

		// Walk the path to configure the summary.
		for i := range f.steps {
			if stepsConditionIndex[i] == 0 {
				// Not found step in the given status's condition list.
				continue
			}
			c := &st.Conditions[stepsConditionIndex[i]-1]

			// Get summary from display result.
			s.Phase, s.PhaseMessage, s.Score = f.stepsDecide[i](c.Status, c.Reason)
			if s.PhaseMessage == "" {
				s.PhaseMessage = c.Message
			}

			// Quit from the walk if still error or being transitioning.
			if s.Score == SummaryScoreInterrupted || s.Score == SummaryScoreTransitioning {
				break
			}
		}
	}

	// Default summary if it hasn't been configured.
	if s == nil || s.Phase == "" {
		s.Phase, s.PhaseMessage, s.Score = f.stepsDecide[len(f.steps)-1]("", "")
	}

	return s
}

// Decision exposes ability to customize how to make a decision on one specified step.
type Decision[T ~string] path[T]

// Make makes a decision on the given specified step with dedicated decide logic.
func (d Decision[T]) Make(step T, with Decide) Decision[T] {
	if with != nil {
		if idx, exist := d.stepsIndex[step]; exist {
			d.stepsDecide[idx] = with
		}
	}

	return d
}

// getGeneralDecide returns a decision that adapts general scene, including,
//   - displays step pretty,
//   - marks step as interrupted if status is False,
//   - marks step as transitioning if status is Unknown,
//   - and moves to next step if status is True.
func getGeneralDecide[T ~string](step T) Decide {
	s := string(step)

	// Pretty the display with some rules,
	// most rules are for not present tense word.
	displays := [3]string{s, s, s} // Transitioning, Error, Done.

	for m, r := range replacements {
		if !strings.HasSuffix(s, m) {
			continue
		}
		p := s[:len(s)-len(m)]
		displays[0], displays[1], displays[2] = p+r.T, p+r.E, p+r.D
	}

	return func(st walruscore.ConditionStatus, _ string) (string, string, Score) {
		switch st {
		case walruscore.ConditionUnknown:
			return displays[0], "", SummaryScoreTransitioning
		case walruscore.ConditionFalse:
			return displays[1], "", SummaryScoreInterrupted
		}

		return displays[2], "", SummaryScoreDone
	}
}

// replacements collects the rules for replacing phased descriptor of the key,
// includes transitioning(T), error(E) and done(D).
var replacements = map[string]struct {
	T, E, D string
}{
	"Running":     {"Running", "Failed", "Completed"},
	"Pending":     {"Pending", "Failed", "Pending"},
	"Progressing": {"Progressing", "Progressing", "Progressed"},
	"Connected":   {"Connecting", "Disconnected", "Connected"},
	"Initialized": {"Initializing", "InitializeFailed", "Initialized"},
	"Scheduled":   {"Scheduling", "ScheduleFailed", "Scheduled"},
	"Accepted":    {"Accepting", "NotAccepted", "Accepted"},
	"Deployed":    {"Deploying", "DeployFailed", "Deployed"},
	"Stopped":     {"Stopping", "StopFailed", "Stopped"},
	"Synced":      {"Syncing", "SyncFailed", "Synced"},
	"Available":   {"Preparing", "Unavailable", "Available"},
	"Ready":       {"Preparing", "NotReady", "Ready"},
	"Active":      {"Preparing", "Inactive", "Active"},
	"Canceled":    {"Canceling", "CancelFailed", "Canceled"},
	"Planned":     {"Planning", "Failed", "Planned"},
	"Applied":     {"Running", "Failed", "Succeeded"},
}
