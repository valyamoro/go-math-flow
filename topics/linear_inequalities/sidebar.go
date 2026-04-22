package ineq

import (
	"fmt"
	"go-math-flow/core"
	"math"
	"strings"
)

func BuildSidebar(ip InequalityProblem, is InequalitySolution) core.SidebarData {
	var standard string
	if ip.TwoVar {
		standard = buildStdForm2Var(ip.A, ip.B, ip.Op)
	} else {
		standard = buildStdForm(ip.A, ip.B, ip.Op)
	}
	sd := core.SidebarData{
		Original: ip.Original(),
		Standard: standard,
		Solution: is.Describe(),
	}
	switch is.SolutionKind() {
	case core.SolHalfPlane:
		sd.ResultKind = "halfplane"
	case core.SolInterval:
		sd.ResultKind = "interval"
		sd.StepMul = buildStepMul(ip.A, ip.B, ip.Op)
		sd.StepDiv = buildStepDiv(ip.A, ip.B, ip.Op)
	case core.SolNone:
		sd.ResultKind = "none"
	case core.SolInfinite:
		sd.ResultKind = "infinite"
	}
	return sd
}

func buildStdForm2Var(A, B float64, op string) string {
	var sb strings.Builder
	sb.WriteString("y")
	const eps = 1e-12
	switch {
	case A > eps:
		sb.WriteString(" - " + fmtS(A) + "x")
	case A < -eps:
		sb.WriteString(" + " + fmtS(-A) + "x")
	}
	switch {
	case B > eps:
		sb.WriteString(" - " + fmtS(B))
	case B < -eps:
		sb.WriteString(" + " + fmtS(-B))
	}
	sb.WriteString(" " + op + " 0")
	return sb.String()
}

func buildStdForm(A, B float64, op string) string {
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
	sb.WriteString(" " + op + " 0")
	return sb.String()
}

func buildStepMul(A, B float64, op string) string {
	var sb strings.Builder
	switch {
	case A == 1:
		sb.WriteString("x")
	case A == -1:
		sb.WriteString("-x")
	default:
		sb.WriteString(fmtS(A) + "x")
	}
	sb.WriteString(" " + op + " " + fmtS(-B))
	return sb.String()
}

func buildStepDiv(A, B float64, op string) string {
	bound := math.Round((-B/A)*1000) / 1000
	num, denom := -B, A
	effectiveOp := op
	if denom < 0 {
		num, denom = -num, -denom
		effectiveOp = flipOp(op)
	}
	return fmt.Sprintf("x %s %s / %s = %s", effectiveOp, fmtS(num), fmtS(denom), fmtS(bound))
}

func flipOp(op string) string {
	switch op {
	case "<":
		return ">"
	case ">":
		return "<"
	case "<=":
		return ">="
	case ">=":
		return "<="
	}
	return op
}

func fmtS(v float64) string {
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
