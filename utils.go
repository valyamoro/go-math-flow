package main

import (
	"fmt"
	"math"
	"strings"
)

func fmtF(v float64) string {
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}

func buildStandardForm(A, B float64) string {
	var sb strings.Builder
	switch {
	case A == 1:
		sb.WriteString("x")
	case A == -1:
		sb.WriteString("-x")
	default:
		sb.WriteString(fmtF(A) + "x")
	}
	const eps = 1e-12
	if B > eps {
		sb.WriteString(" + " + fmtF(B))
	} else if B < -eps {
		sb.WriteString(" - " + fmtF(-B))
	}
	sb.WriteString(" = 0")
	return sb.String()
}

func buildFuncStr(A, B float64) string {
	var sb strings.Builder
	const eps = 1e-12
	if math.Abs(A) < eps {
		return fmtF(B)
	}
	switch {
	case A == 1:
		sb.WriteString("x")
	case A == -1:
		sb.WriteString("-x")
	default:
		sb.WriteString(fmtF(A) + "x")
	}
	if B > eps {
		sb.WriteString(" + " + fmtF(B))
	} else if B < -eps {
		sb.WriteString(" - " + fmtF(-B))
	}
	return sb.String()
}

func buildStepMul(A, B float64) string {
	negB := -B
	lhs := ""
	switch {
	case A == 1:
		lhs = "x"
	case A == -1:
		lhs = "-x"
	default:
		lhs = fmtF(A) + "x"
	}
	return lhs + " = " + fmtF(negB)
}

// r3 rounds v to 3 decimal places — avoids ugly floats in displayed results.
func r3(v float64) float64 { return math.Round(v*1000) / 1000 }

func buildStepDiv(A, B float64) string {
	root := r3(-B / A)
	num, denom := -B, A
	if denom < 0 {
		num, denom = -num, -denom
	}
	return fmt.Sprintf("x = %s / %s = %s", fmtF(num), fmtF(denom), fmtF(root))
}
