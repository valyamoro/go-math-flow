package main

import (
	"html/template"
)

type SolutionKind int

const (
	SolutionUnique   SolutionKind = iota
	SolutionNone
	SolutionInfinite
)

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
