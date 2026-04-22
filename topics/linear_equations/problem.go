package linear

import (
	"go-math-flow/core"
	"math"
)

type LinearProblem struct {
	A        float64
	B        float64
	Op       string
	TwoVar   bool
	original string
}

func New(netX, netY, netConst float64, op, original string) LinearProblem {
	if op == "" {
		op = "="
	}
	if math.Abs(netY) > 1e-12 {
		return LinearProblem{
			A:        -netX / netY,
			B:        -netConst / netY,
			Op:       op,
			TwoVar:   true,
			original: original,
		}
	}
	return LinearProblem{A: netX, B: netConst, Op: op, original: original}
}

func (p LinearProblem) Kind() core.ProblemKind {
	if p.Op == "=" {
		return core.KindLinearEquation
	}
	return core.KindLinearInequality
}

func (p LinearProblem) Original() string { return p.original }
