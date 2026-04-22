package core

type MathProblem interface {
	Kind() ProblemKind
	Original() string
}

type Solution interface {
	SolutionKind() SolutionKind
	Describe() string
}

type Solver interface {
	Accepts(p MathProblem) bool
	Solve(p MathProblem) (Solution, error)
}

type Visualizer interface {
	Accepts(p MathProblem, s Solution) bool
	Render(p MathProblem, s Solution, vp Viewport) (RenderData, error)
}

type Viewport struct {
	XMin, XMax float64
	YMin, YMax float64
}

type RenderData struct {
	TracesJSON string
	LayoutJSON string
	StepHTML   string
}
