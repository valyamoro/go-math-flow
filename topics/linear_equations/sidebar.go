package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
	"strings"
)

type SidebarData struct {
	Original   string
	Standard   string
	Solution   string
	ResultKind string
	Root       float64
	StepMul    string
	StepDiv    string
}

func BuildSidebar(lp LinearProblem, ls LinearSolution) SidebarData {
	sd := SidebarData{
		Original: lp.Original(),
		Standard: buildStdForm(lp.A, lp.B),
		Solution: ls.Describe(),
	}
	switch ls.SolutionKind() {
	case core.SolUnique:
		sd.ResultKind = "unique"
		sd.Root = ls.Root()
		sd.StepMul = buildStpMul(lp.A, lp.B)
		sd.StepDiv = buildStpDiv(lp.A, lp.B)
	case core.SolNone:
		sd.ResultKind = "none"
	case core.SolInfinite:
		sd.ResultKind = "infinite"
	case core.SolLine:
		sd.ResultKind = "line"
	}
	return sd
}

func buildStdForm(A, B float64) string {
	var sb strings.Builder
	const eps = 1e-12
	switch {
	case A == 1:
		sb.WriteString("x")
	case A == -1:
		sb.WriteString("-x")
	default:
		sb.WriteString(fmtS(A) + "x")
	}
	if B > eps {
		sb.WriteString(" + " + fmtS(B))
	} else if B < -eps {
		sb.WriteString(" - " + fmtS(-B))
	}
	sb.WriteString(" = 0")
	return sb.String()
}

func buildStpMul(A, B float64) string {
	lhs := ""
	switch {
	case A == 1:
		lhs = "x"
	case A == -1:
		lhs = "-x"
	default:
		lhs = fmtS(A) + "x"
	}
	return lhs + " = " + fmtS(-B)
}

func buildStpDiv(A, B float64) string {
	root := math.Round((-B/A)*1000) / 1000
	num, denom := -B, A
	if denom < 0 {
		num, denom = -num, -denom
	}
	return fmt.Sprintf("x = %s / %s = %s", fmtS(num), fmtS(denom), fmtS(root))
}

func fmtS(v float64) string {
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
