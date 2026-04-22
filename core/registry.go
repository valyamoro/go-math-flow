package core

import "fmt"

var (
	solvers     []Solver
	visualizers []Visualizer
)

// RegisterSolver adds a Solver to the global registry.
// Call this inside a topic package's init() so it auto-registers on import.
//
//	func init() { core.RegisterSolver(MyTopicSolver{}) }
func RegisterSolver(s Solver) {
	solvers = append(solvers, s)
}

// RegisterVisualizer adds a Visualizer to the global registry.
// Call this inside a viz package's init() so it auto-registers on import.
func RegisterVisualizer(v Visualizer) {
	visualizers = append(visualizers, v)
}

// Solve walks the solver registry and delegates to the first one that Accepts p.
// Returns an error if no solver is registered for the problem's kind.
func Solve(p MathProblem) (Solution, error) {
	for _, s := range solvers {
		if s.Accepts(p) {
			return s.Solve(p)
		}
	}
	return nil, fmt.Errorf("no solver registered for problem kind %d", p.Kind())
}

// Render walks the visualizer registry and delegates to the first one that Accepts (p, s).
// Returns an error if no visualizer is registered for the combination.
func Render(p MathProblem, s Solution, vp Viewport) (RenderData, error) {
	for _, v := range visualizers {
		if v.Accepts(p, s) {
			return v.Render(p, s, vp)
		}
	}
	return RenderData{}, fmt.Errorf("no visualizer registered for problem kind %d", p.Kind())
}
