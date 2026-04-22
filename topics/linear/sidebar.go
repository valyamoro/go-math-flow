package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
	"strings"
)

// SidebarData holds all display-ready strings for the HTML sidebar panel.
// It is built once from LinearProblem + LinearSolution and passed directly
// to the template — the template never touches math logic.
type SidebarData struct {
	Original   string  // raw input:           "2x + 3 = 7"
	Standard   string  // normal form:         "2x - 4 = 0" / "2x - 4 > 0"
	Solution   string  // human answer:        "x = 2" / "x > 3" / "x ∈ ℝ"
	ResultKind string  // "unique" | "none" | "infinite" | "interval"
	Root       float64 // numeric root — only when ResultKind == "unique"
	StepMul    string  // step 2: "2x = 4"    — only when ResultKind == "unique"
	StepDiv    string  // step 3: "x = 4/2=2" — only when ResultKind == "unique"
}

// BuildSidebar assembles a SidebarData ready for the HTML template.
func BuildSidebar(lp LinearProblem, ls LinearSolution) SidebarData {
	sd := SidebarData{
		Original: lp.Original(),
		Standard: buildStdForm(lp.A, lp.B, lp.Op),
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
	case core.SolInterval:
		sd.ResultKind = "interval"
	}
	return sd
}

// buildStdForm renders the canonical "Ax + B ⋈ 0" string.
// Using the actual operator (not always "=") fixes the display for inequalities.
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

// fmtS formats a float without trailing zeros. Private to this package.
func fmtS(v float64) string {
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
