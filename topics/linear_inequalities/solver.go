package ineq

import (
	"fmt"
	"go-math-flow/core"
	"math"
)

type InequalitySolver struct{}

func (InequalitySolver) Accepts(p core.MathProblem) bool {
	return p.Kind() == core.KindLinearInequality
}

func (InequalitySolver) Solve(p core.MathProblem) (core.Solution, error) {
	ip, ok := p.(InequalityProblem)
	if !ok {
		return nil, fmt.Errorf("InequalitySolver: expected InequalityProblem, got %T", p)
	}

	const eps = 1e-12
	if ip.TwoVar {
		positive := ip.Op == ">" || ip.Op == ">="
		strict := ip.Op == "<" || ip.Op == ">"
		return InequalitySolution{
			kind:      core.SolHalfPlane,
			slope:     ip.A,
			intercept: ip.B,
			op:        ip.Op,
			positive:  positive,
			strict:    strict,
		}, nil
	}
	if math.Abs(ip.A) < eps {
		if evalConst(ip.B, ip.Op) {
			return InequalitySolution{kind: core.SolInfinite, op: ip.Op}, nil
		}
		return InequalitySolution{kind: core.SolNone, op: ip.Op}, nil
	}

	bound := math.Round((-ip.B/ip.A)*1000) / 1000
	positive, strict := intervalDirection(ip.Op, ip.A < 0)
	return InequalitySolution{
		kind:     core.SolInterval,
		bound:    bound,
		op:       ip.Op,
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
		return b < 0
	case ">":
		return b > eps
	case "<=":
		return b <= eps
	case ">=":
		return b >= -eps
	}
	return false
}
