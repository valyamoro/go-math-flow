package main

import (
	"flag"
	"fmt"
	"html/template"
	"math"
	"os"
	"path/filepath"

	"go-math-flow/core"
	"go-math-flow/parser"
	"go-math-flow/topics/linear"
	_ "go-math-flow/viz/cartesian" // registers CartesianVisualizer via init()
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

	flag.StringVar(&rawEq, "i", "", "Equation or inequality. Example: -i '2x + 3 = 7'")
	flag.StringVar(&out, "o", "tmp/equation.html", "Output HTML file")
	flag.Float64Var(&xmin, "xmin", -10, "Min X")
	flag.Float64Var(&xmax, "xmax", 10, "Max X")
	flag.Float64Var(&ymin, "ymin", -10, "Min Y")
	flag.Float64Var(&ymax, "ymax", 10, "Max Y")
	flag.Parse()

	if rawEq == "" {
		rawEq = "2x + 3 = 7"
		fmt.Println("Input not specified — using built-in example:", rawEq)
	}

	// --- new pipeline ---
	problem, err := parser.Parse(rawEq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	lp := problem.(linear.LinearProblem)
	fmt.Printf("Input: %s  →  A=%-6g B=%-6g op=%s\n", rawEq, lp.A, lp.B, lp.Op)

	// Auto-adjust Y viewport for degenerate case (A ≈ 0 → horizontal line)
	if math.Abs(lp.A) < 1e-12 {
		ymin = lp.B - 5
		ymax = lp.B + 5
	}

	solution, err := core.Solve(problem)
	if err != nil {
		fmt.Fprintf(os.Stderr, "solve error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Solution:", solution.Describe())

	vp := core.Viewport{XMin: xmin, XMax: xmax, YMin: ymin, YMax: ymax}
	rd, err := core.Render(problem, solution, vp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "render error: %v\n", err)
		os.Exit(1)
	}

	// Build PageData for the existing HTML template
	pd := newPageData(lp, solution, rd, vp)
	if err := renderHTML(pd, out); err != nil {
		fmt.Fprintf(os.Stderr, "html error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done →", out)
}

// newPageData bridges core.RenderData + LinearSolution into the legacy PageData
// that the HTML template expects. This adapter will shrink once the template is
// migrated in a later step.
func newPageData(lp linear.LinearProblem, sol core.Solution, rd core.RenderData, vp core.Viewport) PageData {
	ls := sol.(linear.LinearSolution)

	ei := EqInfo{
		Original: lp.Original(),
		Standard: buildStandardForm(lp.A, lp.B),
		FuncStr:  buildFuncStr(lp.A, lp.B),
		Kind:     legacyKind(ls.SolutionKind()),
	}
	if ls.SolutionKind() == core.SolUnique {
		ei.Root = ls.Root()
		ei.HasRoot = true
		ei.StepMul = buildStepMul(lp.A, lp.B)
		ei.StepDiv = buildStepDiv(lp.A, lp.B)
	}

	return PageData{
		TracesJSON: template.JS(rd.TracesJSON),
		LayoutJSON: template.JS(rd.LayoutJSON),
		Eq:         ei,
		XMin:       vp.XMin, XMax: vp.XMax,
		YMin: vp.YMin, YMax: vp.YMax,
	}
}

// legacyKind maps core.SolutionKind back to the old SolutionKind enum so the
// existing HTML template keeps working without changes.
func legacyKind(k core.SolutionKind) SolutionKind {
	switch k {
	case core.SolUnique:
		return SolutionUnique
	case core.SolNone:
		return SolutionNone
	default:
		return SolutionInfinite
	}
}
