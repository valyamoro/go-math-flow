package linear

import (
	"go-math-flow/core"
	"math"
)

type LinearProblem struct {
	A        float64
	B        float64
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
			TwoVar:   true,
			original: original,
		}
	}
	return LinearProblem{A: netX, B: netConst, original: original}
}

func (p LinearProblem) Kind() core.ProblemKind {
	return core.KindLinearEquation
}

func (p LinearProblem) Original() string { return p.original }
