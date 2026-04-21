package main

import (
	"html/template"
	"strings"
)

type SolutionKind int

const (
	SolutionUnique   SolutionKind = iota
	SolutionNone
	SolutionInfinite
)

type Equation struct {
	A, B     float64
	Original string
}

type EqInfo struct {
	Original string
	Standard string
	FuncStr  string
	Kind     SolutionKind
	Root     float64
	HasRoot  bool
	StepMul  string
	StepDiv  string
}

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
