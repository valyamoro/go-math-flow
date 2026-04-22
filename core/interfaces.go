package core

// MathProblem is any mathematical problem that can be parsed and solved.
// Every topic (linear eq, quadratic, system, integral...) implements this.
type MathProblem interface {
	Kind()     ProblemKind
	Original() string
}

// Solution is the result of solving a MathProblem.
// The shape of the answer depends on SolutionKind: a point, interval, vector, etc.
type Solution interface {
	SolutionKind() SolutionKind
	Describe() string
}

// Solver resolves a specific category of MathProblem.
// Each topic registers its own Solver via RegisterSolver in init().
type Solver interface {
	Accepts(p MathProblem) bool
	Solve(p MathProblem) (Solution, error)
}

// Visualizer turns a solved problem into raw JSON data for the HTML template.
// Decoupled from Solver: one problem type can have multiple visualizations,
// and one visualizer can serve multiple problem types.
type Visualizer interface {
	Accepts(p MathProblem, s Solution) bool
	Render(p MathProblem, s Solution, vp Viewport) (RenderData, error)
}

// Viewport defines the coordinate window for 2D graphs.
type Viewport struct {
	XMin, XMax float64
	YMin, YMax float64
}

// RenderData is the output of a Visualizer — raw JSON ready for the HTML template.
// StepHTML carries optional step-by-step explanation markup.
type RenderData struct {
	TracesJSON string
	LayoutJSON string
	StepHTML   string
}
