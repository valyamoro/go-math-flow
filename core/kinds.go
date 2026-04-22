package core

// ProblemKind identifies the mathematical category of a MathProblem.
// New topics are added here — nothing else changes in the core.
type ProblemKind int

const (
	KindLinearEquation   ProblemKind = iota // ax + b = 0
	KindLinearInequality                    // ax + b < 0, >, <=, >=
	KindQuadraticEquation                   // ax² + bx + c = 0
	KindLinearSystem                        // system of linear equations
	KindDerivative                          // f'(x)
	KindIntegral                            // ∫f(x)dx
)

// SolutionKind describes the shape of the answer set.
// Solvers set this so Visualizers know how to draw the result.
type SolutionKind int

const (
	SolUnique   SolutionKind = iota // single value:        x = k
	SolNone                         // empty set:           no solution
	SolInfinite                     // all real numbers:    x ∈ ℝ
	SolInterval                     // half-line or segment: x > k, a ≤ x ≤ b
	SolSet                          // finite discrete set: {x₁, x₂, ...}
	SolVector                       // solution is a vector (linear algebra)
)
