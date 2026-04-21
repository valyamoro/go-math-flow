package main

import "math"

func (eq Equation) Solve() (SolutionKind, float64) {
	const eps = 1e-12
	if math.Abs(eq.A) > eps {
		return SolutionUnique, -eq.B / eq.A
	}
	if math.Abs(eq.B) > eps {
		return SolutionNone, 0
	}
	return SolutionInfinite, 0
}

func r3(v float64) float64 { return math.Round(v*1000) / 1000 }
