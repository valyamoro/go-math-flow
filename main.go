package main

import (
	"flag"
	"fmt"
	"html/template"
	"math"
	"os"
	"path/filepath"
)

var tmplFuncs = template.FuncMap{
	"fmtF":     fmtF,
	"isUnique": func(k SolutionKind) bool { return k == SolutionUnique },
	"isNone":   func(k SolutionKind) bool { return k == SolutionNone },
}

func renderHTML(pd PageData, outPath string) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}
	tmpl, err := template.New("page").Funcs(tmplFuncs).Parse(htmlTemplate)
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, pd)
}

func main() {
	var rawEq string
	var out string
	var xmin, xmax, ymin, ymax float64

	flag.StringVar(&rawEq, "i", "", "Equation. Example: -i '2x + 3 = 7'")
	flag.StringVar(&out, "o", "tmp/equation.html", "Output HTML file")
	flag.Float64Var(&xmin, "xmin", -10, "Min X")
	flag.Float64Var(&xmax, "xmax", 10, "Max X")
	flag.Float64Var(&ymin, "ymin", -10, "Min Y")
	flag.Float64Var(&ymax, "ymax", 10, "Max Y")
	flag.Parse()

	if rawEq == "" {
		rawEq = "2x + 3 = 7"
		fmt.Println("Equation not specified — using built-in example:", rawEq)
	}

	eq, err := ParseEquation(rawEq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Equation: %s  →  A=%-6g B=%-6g\n", rawEq, eq.A, eq.B)

	// Авто-масштаб Y если уравнение только константа
	if math.Abs(eq.A) < 1e-12 {
		ymin = eq.B - 5
		ymax = eq.B + 5
	}

	pd, err := buildPageData(eq, xmin, xmax, ymin, ymax)
	if err != nil {
		fmt.Fprintf(os.Stderr, "build error: %v\n", err)
		os.Exit(1)
	}
	if err := renderHTML(pd, out); err != nil {
		fmt.Fprintf(os.Stderr, "render error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done →", out)
}
