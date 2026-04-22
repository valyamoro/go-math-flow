package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

// LinearSolution is the result of solving a LinearProblem.
// For equations:    SolUnique (x = k) | SolNone | SolInfinite
// For inequalities: SolInterval (x > k, x ≤ k, etc.) | SolNone | SolInfinite
type LinearSolution struct {
	kind     core.SolutionKind
	root     float64 // valid when kind == SolUnique
	bound    float64 // valid when kind == SolInterval (the boundary value)
	op       string  // the original relation operator
	strict   bool    // true for < / >, false for <= / >=
	positive bool    // true if the interval is x > bound (or x >= bound)
}

func (s LinearSolution) SolutionKind() core.SolutionKind { return s.kind }

func (s LinearSolution) Describe() string {
	switch s.kind {
	case core.SolUnique:
		return fmt.Sprintf("x = %s", fmtV(s.root))
	case core.SolNone:
		return "no solution (∅)"
	case core.SolInfinite:
		return "x ∈ ℝ"
	case core.SolInterval:
		if s.positive {
			if s.strict {
				return fmt.Sprintf("x > %s", fmtV(s.bound))
			}
			return fmt.Sprintf("x ≥ %s", fmtV(s.bound))
		}
		if s.strict {
			return fmt.Sprintf("x < %s", fmtV(s.bound))
		}
		return fmt.Sprintf("x ≤ %s", fmtV(s.bound))
	}
	return ""
}

// Root returns the unique solution value. Only meaningful when SolutionKind == SolUnique.
func (s LinearSolution) Root() float64 { return s.root }

// Bound returns the boundary of the interval solution. Only meaningful when SolutionKind == SolInterval.
func (s LinearSolution) Bound() float64 { return s.bound }

// IsPositiveDirection reports whether the interval opens to +∞ (x > k or x ≥ k).
func (s LinearSolution) IsPositiveDirection() bool { return s.positive }

// IsStrict reports whether the boundary is excluded (< or >).
func (s LinearSolution) IsStrict() bool { return s.strict }

func fmtV(v float64) string {
	v = math.Round(v*1000) / 1000
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
