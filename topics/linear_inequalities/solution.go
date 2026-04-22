package ineq

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

type InequalitySolution struct {
	kind     core.SolutionKind
	bound    float64
	op       string
	strict   bool
	positive bool
}

func (s InequalitySolution) SolutionKind() core.SolutionKind { return s.kind }

func (s InequalitySolution) Describe() string {
	switch s.kind {
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

func (s InequalitySolution) Bound() float64            { return s.bound }
func (s InequalitySolution) IsPositiveDirection() bool { return s.positive }
func (s InequalitySolution) IsStrict() bool            { return s.strict }

func fmtV(v float64) string {
	v = math.Round(v*1000) / 1000
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
