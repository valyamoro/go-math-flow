package core

import (
	"fmt"
	"math"
	"strings"
)

// Expr is a node in a mathematical expression tree (AST).
// This is the foundation for every formula in the system — from "2x+3"
// to "sin(x²) + ln(x)". All parsers produce an Expr; all solvers consume it.
type Expr interface {
	// Eval computes the numeric value given variable bindings.
	Eval(vars map[string]float64) float64
	// String returns a human-readable representation.
	String() string
}

// Num is a numeric constant leaf node, e.g. 3.14
type Num struct{ V float64 }

func (n Num) Eval(_ map[string]float64) float64 { return n.V }
func (n Num) String() string {
	if n.V == math.Trunc(n.V) {
		return fmt.Sprintf("%.0f", n.V)
	}
	return fmt.Sprintf("%g", n.V)
}

// Var is a variable leaf node, e.g. "x", "y", "t"
type Var struct{ Name string }

func (v Var) Eval(vars map[string]float64) float64 { return vars[v.Name] }
func (v Var) String() string                        { return v.Name }

// BinOp is a binary operation: +  -  *  /  ^
type BinOp struct {
	Op   rune
	L, R Expr
}

func (b BinOp) Eval(vars map[string]float64) float64 {
	l, r := b.L.Eval(vars), b.R.Eval(vars)
	switch b.Op {
	case '+':
		return l + r
	case '-':
		return l - r
	case '*':
		return l * r
	case '/':
		if r == 0 {
			return math.NaN()
		}
		return l / r
	case '^':
		return math.Pow(l, r)
	}
	return 0
}

func (b BinOp) String() string {
	return fmt.Sprintf("(%s %c %s)", b.L, b.Op, b.R)
}

// UnaryOp is a unary function or negation: -, sin, cos, tan, ln, log, sqrt, abs
type UnaryOp struct {
	Op  string
	Arg Expr
}

func (u UnaryOp) Eval(vars map[string]float64) float64 {
	a := u.Arg.Eval(vars)
	switch strings.ToLower(u.Op) {
	case "-":
		return -a
	case "sin":
		return math.Sin(a)
	case "cos":
		return math.Cos(a)
	case "tan":
		return math.Tan(a)
	case "ln":
		return math.Log(a)
	case "log":
		return math.Log10(a)
	case "sqrt":
		return math.Sqrt(a)
	case "abs":
		return math.Abs(a)
	}
	return 0
}

func (u UnaryOp) String() string {
	if u.Op == "-" {
		return fmt.Sprintf("-%s", u.Arg)
	}
	return fmt.Sprintf("%s(%s)", u.Op, u.Arg)
}
