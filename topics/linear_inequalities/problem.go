package ineq

import (
	"go-math-flow/core"
	"math"
)

type InequalityProblem struct {
	A        float64
	B        float64
	Op       string
	TwoVar   bool
	original string
}

func New(netX, netY, netConst float64, op, original string) InequalityProblem {
	if math.Abs(netY) > 1e-12 {
		effectiveOp := op
		if netY < 0 {
			effectiveOp = flipOp(op)
		}
		return InequalityProblem{
			A:        -netX / netY,
			B:        -netConst / netY,
			Op:       effectiveOp,
			TwoVar:   true,
			original: original,
		}
	}
	return InequalityProblem{A: netX, B: netConst, Op: op, original: original}
}

func (p InequalityProblem) Kind() core.ProblemKind { return core.KindLinearInequality }
func (p InequalityProblem) Original() string       { return p.original }
