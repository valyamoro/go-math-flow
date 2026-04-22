package core

type ProblemKind int

const (
	KindLinearEquation ProblemKind = iota
	KindLinearInequality
	KindQuadraticEquation
	KindLinearSystem
	KindDerivative
	KindIntegral
)

type SolutionKind int

const (
	SolUnique SolutionKind = iota
	SolNone
	SolInfinite
	SolInterval
	SolSet
	SolVector
	SolLine      // 2-var: y = Ax + B
	SolHalfPlane // 2-var: y op Ax + B
)
