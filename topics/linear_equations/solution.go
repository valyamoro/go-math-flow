package equat

import (
	"fmt"
	"go-math-flow/core"
	"math"
	"strings"
)

type LinearSolution struct {
	kind      core.SolutionKind
	root      float64
	slope     float64
	intercept float64
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
	case core.SolLine:
		return buildSlopeInterceptStr(s.slope, s.intercept)
	}
	return ""
}

func (s LinearSolution) Root() float64      { return s.root }
func (s LinearSolution) Slope() float64     { return s.slope }
func (s LinearSolution) Intercept() float64 { return s.intercept }

func buildSlopeInterceptStr(a, b float64) string {
	const eps = 1e-12
	var sb strings.Builder
	sb.WriteString("y = ")
	switch {
	case math.Abs(a) < eps:
		sb.WriteString(fmtV(b))
		return sb.String()
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
	return sb.String()
}

func fmtV(v float64) string {
	v = math.Round(v*1000) / 1000
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
