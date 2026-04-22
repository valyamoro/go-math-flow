package ineq

import "go-math-flow/core"

type InequalityProblem struct {
	A        float64
	B        float64
	Op       string
	original string
}

func New(netX, netConst float64, op, original string) InequalityProblem {
	return InequalityProblem{A: netX, B: netConst, Op: op, original: original}
}

func (p InequalityProblem) Kind() core.ProblemKind    { return core.KindLinearInequality }
func (p InequalityProblem) Original() string          { return p.original }
