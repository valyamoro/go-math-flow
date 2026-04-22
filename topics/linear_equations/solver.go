package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

type LinearSolver struct{}

func (LinearSolver) Accepts(p core.MathProblem) bool {
	return p.Kind() == core.KindLinearEquation
}

func (LinearSolver) Solve(p core.MathProblem) (core.Solution, error) {
	lp, ok := p.(LinearProblem)
	if !ok {
		return nil, fmt.Errorf("LinearSolver: expected LinearProblem, got %T", p)
	}

	if lp.TwoVar {
		return LinearSolution{kind: core.SolLine, slope: lp.A, intercept: lp.B}, nil
	}

	const eps = 1e-12
	if math.Abs(lp.A) < eps {
		if math.Abs(lp.B) < eps {
			return LinearSolution{kind: core.SolInfinite}, nil
		}
		return LinearSolution{kind: core.SolNone}, nil
	}

	return LinearSolution{kind: core.SolUnique, root: -lp.B / lp.A}, nil
}
