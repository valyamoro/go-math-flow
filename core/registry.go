package core

import "fmt"

var (
	solvers     []Solver
	visualizers []Visualizer
)

func RegisterSolver(s Solver) {
	solvers = append(solvers, s)
}

func RegisterVisualizer(v Visualizer) {
	visualizers = append(visualizers, v)
}

func Solve(p MathProblem) (Solution, error) {
	for _, s := range solvers {
		if s.Accepts(p) {
			return s.Solve(p)
		}
	}
	return nil, fmt.Errorf("no solver registered for problem kind %d", p.Kind())
}

func Render(p MathProblem, s Solution, vp Viewport) (RenderData, error) {
	for _, v := range visualizers {
		if v.Accepts(p, s) {
			return v.Render(p, s, vp)
		}
	}
	return RenderData{}, fmt.Errorf("no visualizer registered for problem kind %d", p.Kind())
}
