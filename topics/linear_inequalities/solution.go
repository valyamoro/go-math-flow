package ineq

import (
	"fmt"
	"go-math-flow/core"
	"math"
	"strings"
)

type InequalitySolution struct {
	kind      core.SolutionKind
	bound     float64
	slope     float64
	intercept float64
	op        string
	strict    bool
	positive  bool
}

func (s InequalitySolution) SolutionKind() core.SolutionKind { return s.kind }

func (s InequalitySolution) Describe() string {
	switch s.kind {
	case core.SolNone:
		return "no solution (∅)"
	case core.SolInfinite:
		return "x ∈ ℝ"
	case core.SolHalfPlane:
		return buildHalfPlaneStr(s.slope, s.intercept, s.op)
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
func (s InequalitySolution) Slope() float64            { return s.slope }
func (s InequalitySolution) Intercept() float64        { return s.intercept }
func (s InequalitySolution) IsPositiveDirection() bool { return s.positive }
func (s InequalitySolution) IsStrict() bool            { return s.strict }

func buildHalfPlaneStr(a, b float64, op string) string {
	const eps = 1e-12
	var sb strings.Builder
	sb.WriteString("y " + op + " ")
	hasX := math.Abs(a) > eps
	if hasX {
		switch {
		case a == 1:
			sb.WriteString("x")
		case a == -1:
			sb.WriteString("-x")
		default:
			sb.WriteString(fmtV(a) + "x")
		}
		if b > eps {
			sb.WriteString(" + " + fmtV(b))
		} else if b < -eps {
			sb.WriteString(" - " + fmtV(-b))
		}
	} else {
		sb.WriteString(fmtV(b))
	}
	return sb.String()
}

func fmtV(v float64) string {
	v = math.Round(v*1000) / 1000
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
