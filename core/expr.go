package core

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(vars map[string]float64) float64
	String() string
}

type Num struct{ V float64 }

func (n Num) Eval(_ map[string]float64) float64 { return n.V }
func (n Num) String() string {
	if n.V == math.Trunc(n.V) {
		return fmt.Sprintf("%.0f", n.V)
	}
	return fmt.Sprintf("%g", n.V)
}

type Var struct{ Name string }

func (v Var) Eval(vars map[string]float64) float64 { return vars[v.Name] }
func (v Var) String() string                       { return v.Name }

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
