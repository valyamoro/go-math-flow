package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

type LinearSolution struct {
	kind     core.SolutionKind
	root     float64
	bound    float64
	op       string
	strict   bool
	positive bool
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

func (s LinearSolution) Root() float64             { return s.root }
func (s LinearSolution) Bound() float64            { return s.bound }
func (s LinearSolution) IsPositiveDirection() bool { return s.positive }
func (s LinearSolution) IsStrict() bool            { return s.strict }

func fmtV(v float64) string {
	v = math.Round(v*1000) / 1000
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
