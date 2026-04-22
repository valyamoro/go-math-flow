package linear

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

// LinearSolver handles both LinearEquation and LinearInequality problems.
// It is registered once via init() and never needs to be touched when
// other topics are added.
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

	// Degenerate case: 0x + B ⋈ 0
	if math.Abs(lp.A) < eps {
		if lp.Op == "=" {
			if math.Abs(lp.B) < eps {
				return LinearSolution{kind: core.SolInfinite, op: lp.Op}, nil
			}
			return LinearSolution{kind: core.SolNone, op: lp.Op}, nil
		}
		// inequality with A=0: check if "0 ⋈ 0" (or "B ⋈ 0") is true
		if evalConst(lp.B, lp.Op) {
			return LinearSolution{kind: core.SolInfinite, op: lp.Op}, nil
		}
		return LinearSolution{kind: core.SolNone, op: lp.Op}, nil
	}

	// Normal case: Ax + B ⋈ 0  →  x ⋈ -B/A
	// IMPORTANT: if A < 0 and op is an inequality, the direction flips.
	bound := -lp.B / lp.A

	if lp.Op == "=" {
		return LinearSolution{
			kind: core.SolUnique,
			root: bound,
			op:   lp.Op,
		}, nil
	}

	// Determine interval direction, accounting for sign flip when A < 0.
	positive, strict := intervalDirection(lp.Op, lp.A < 0)
	return LinearSolution{
		kind:     core.SolInterval,
		bound:    math.Round(bound*1000) / 1000,
		op:       lp.Op,
		positive: positive,
		strict:   strict,
	}, nil
}

// intervalDirection returns (positive, strict) for an interval solution.
// flip is true when we divided by a negative A, which reverses the inequality.
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

// evalConst checks whether "0 ⋈ 0" (constant inequality) holds.
// Used when A == 0.
func evalConst(b float64, op string) bool {
	const eps = 1e-12
	switch op {
	case "<":
		return -b < 0 // 0 + B < 0  →  B < 0? No, we have 0x+B<0 → B<0
	case ">":
		return b > eps
	case "<=":
		return b < eps
	case ">=":
		return b > -eps
	}
	return false
}
