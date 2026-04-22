package linear

import "go-math-flow/core"

type LinearProblem struct {
	A, B     float64
	Op       string
	original string
}

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
