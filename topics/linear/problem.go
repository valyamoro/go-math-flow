package linear

import "go-math-flow/core"

// LinearProblem represents an equation or inequality of the form Ax + B ⋈ 0,
// where ⋈ is "=", "<", ">", "<=", ">=".
// It implements core.MathProblem so the registry can dispatch to LinearSolver.
type LinearProblem struct {
	A, B     float64
	Op       string // "=", "<", ">", "<=", ">=" — defaults to "="
	original string
}

// New creates a LinearProblem from coefficients already reduced to Ax + B ⋈ 0 form.
// original is the raw input string, preserved for display purposes.
func New(a, b float64, op, original string) LinearProblem {
	if op == "" {
		op = "="
	}
	return LinearProblem{A: a, B: b, Op: op, original: original}
}

func (p LinearProblem) Kind() core.ProblemKind {
	if p.Op == "=" {
		return core.KindLinearEquation
	}
	return core.KindLinearInequality
}

func (p LinearProblem) Original() string { return p.original }
