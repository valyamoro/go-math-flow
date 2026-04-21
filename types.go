package main

import (
	"html/template"
	"strings"
)

// SolutionKind — тип решения уравнения Ax + B = 0.
type SolutionKind int

const (
	SolutionUnique   SolutionKind = iota // a ≠ 0: один корень
	SolutionNone                         // a = 0, b ≠ 0: нет корней
	SolutionInfinite                     // a = 0, b = 0: любое число
)

// Equation хранит уравнение в стандартной форме Ax + B = 0.
type Equation struct {
	A, B     float64
	Original string
}

// EqInfo передаётся в HTML-шаблон.
type EqInfo struct {
	Original string
	Standard string // "Ax + B = 0"
	FuncStr  string // "Ax + B" для подписи "y = ..."
	Kind     SolutionKind
	Root     float64
	HasRoot  bool
	StepMul  string // шаг 2: "Ax = -B"
	StepDiv  string // шаг 3: "x = -B/A = value"
}

// PageData — данные страницы.
type PageData struct {
	TracesJSON template.JS
	LayoutJSON template.JS
	Eq         EqInfo
	XMin, XMax float64
	YMin, YMax float64
}

type stringsFlag []string

func (f *stringsFlag) String() string     { return strings.Join(*f, ", ") }
func (f *stringsFlag) Set(v string) error { *f = append(*f, v); return nil }
