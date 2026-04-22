package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

type LinearSolver struct{}

func (LinearSolver) Accepts(p core.MathProblem) bool {
	return p.Kind() == core.KindLinearEquation || p.Kind() == core.KindLinearInequality
}

func (LinearSolver) Solve(p core.MathProblem) (core.Solution, error) {
	lp, ok := p.(LinearProblem)
	if !ok {
		return nil, fmt.Errorf("LinearSolver: expected LinearProblem, got %T", p)
	}

	const eps = 1e-12
	if math.Abs(lp.A) < eps {
		if lp.Op == "=" {
			if math.Abs(lp.B) < eps {
				return LinearSolution{kind: core.SolInfinite, op: lp.Op}, nil
			}
			return LinearSolution{kind: core.SolNone, op: lp.Op}, nil
		}
		if evalConst(lp.B, lp.Op) {
			return LinearSolution{kind: core.SolInfinite, op: lp.Op}, nil
		}
		return LinearSolution{kind: core.SolNone, op: lp.Op}, nil
	}

	bound := -lp.B / lp.A

	if lp.Op == "=" {
		return LinearSolution{
			kind: core.SolUnique,
			root: bound,
			op:   lp.Op,
		}, nil
	}

	positive, strict := intervalDirection(lp.Op, lp.A < 0)
	return LinearSolution{
		kind:     core.SolInterval,
		bound:    math.Round(bound*1000) / 1000,
		op:       lp.Op,
		positive: positive,
		strict:   strict,
	}, nil
}

func intervalDirection(op string, flip bool) (positive, strict bool) {
	switch op {
	case "<":
		positive, strict = false, true
	case ">":
		positive, strict = true, true
	case "<=":
		positive, strict = false, false
	case ">=":
		positive, strict = true, false
	}
	if flip {
		positive = !positive
	}
	return
}

func evalConst(b float64, op string) bool {
	const eps = 1e-12
	switch op {
	case "<":
		return -b < 0
	case ">":
		return b > eps
	case "<=":
		return b < eps
	case ">=":
		return b > -eps
	}
	return false
}
