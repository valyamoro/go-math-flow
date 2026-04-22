package cartesian

import (
	"encoding/json"
	"fmt"
	"go-math-flow/core"
	"go-math-flow/topics/linear"
	"math"
)

type m = map[string]interface{}

type CartesianVisualizer struct{}

func (CartesianVisualizer) Accepts(p core.MathProblem, _ core.Solution) bool {
	return p.Kind() == core.KindLinearEquation || p.Kind() == core.KindLinearInequality
}

func (CartesianVisualizer) Render(p core.MathProblem, s core.Solution, vp core.Viewport) (core.RenderData, error) {
	lp := p.(linear.LinearProblem)
	ls := s.(linear.LinearSolution)

	funcStr := buildFuncStr(lp.A, lp.B)
	traces := buildTraces(lp, ls, funcStr, vp)
	layout := buildLayout(vp)

	tj, err := json.Marshal(traces)
	if err != nil {
		return core.RenderData{}, err
	}
	lj, err := json.Marshal(layout)
	if err != nil {
		return core.RenderData{}, err
	}

	return core.RenderData{
		TracesJSON: string(tj),
		LayoutJSON: string(lj),
	}, nil
}

func buildTraces(lp linear.LinearProblem, ls linear.LinearSolution, funcStr string, vp core.Viewport) []m {
	var traces []m

	const nPts = 300
	xs := make([]float64, nPts)
	ys := make([]float64, nPts)
	step := (vp.XMax - vp.XMin) / float64(nPts-1)
	for i := range xs {
		x := vp.XMin + float64(i)*step
		xs[i] = x
		ys[i] = lp.A*x + lp.B
	}
	traces = append(traces, m{
		"type": "scatter",
		"x":    xs,
		"y":    ys,
		"mode": "lines",
		"name": "y = " + funcStr,
		"line": m{"color": "#4EC9DC", "width": 2.8},
	})

	switch ls.SolutionKind() {
	case core.SolUnique:
		root := ls.Root()
		traces = append(traces, m{
			"type":       "scatter",
			"x":          []float64{root, root},
			"y":          []float64{vp.YMin, vp.YMax},
			"mode":       "lines",
			"showlegend": false,
			"line":       m{"color": "rgba(0,230,160,0.22)", "width": 1, "dash": "dot"},
			"hoverinfo":  "skip",
		})
		label := fmt.Sprintf("x = %s", fmtV(root))
		traces = append(traces, m{
			"type":         "scatter",
			"x":            []float64{root},
			"y":            []float64{0},
			"mode":         "markers+text",
			"text":         []string{label},
			"textposition": "top right",
			"textfont":     m{"color": "#00e6a0", "size": 11, "family": "Fira Code, monospace"},
			"marker":       m{"color": "#00e6a0", "size": 9, "line": m{"color": "#0a1628", "width": 2}},
			"name":         "Root",
			"hoverinfo":    "text",
		})

	case core.SolInterval:
		bound := ls.Bound()
		dash := "solid"
		if ls.IsStrict() {
			dash = "dash"
		}
		traces = append(traces, m{
			"type":       "scatter",
			"x":          []float64{bound, bound},
			"y":          []float64{vp.YMin, vp.YMax},
			"mode":       "lines",
			"showlegend": false,
			"line":       m{"color": "rgba(255,180,0,0.5)", "width": 1.5, "dash": dash},
			"hoverinfo":  "skip",
		})
		var shadeX []float64
		if ls.IsPositiveDirection() {
			shadeX = []float64{bound, vp.XMax, vp.XMax, bound}
		} else {
			shadeX = []float64{vp.XMin, bound, bound, vp.XMin}
		}
		shadeY := []float64{vp.YMin, vp.YMin, vp.YMax, vp.YMax}
		traces = append(traces, m{
			"type":       "scatter",
			"x":          shadeX,
			"y":          shadeY,
			"fill":       "toself",
			"fillcolor":  "rgba(255,180,0,0.07)",
			"mode":       "none",
			"showlegend": false,
			"hoverinfo":  "skip",
		})
		label := fmt.Sprintf("x = %s", fmtV(bound))
		traces = append(traces, m{
			"type":         "scatter",
			"x":            []float64{bound},
			"y":            []float64{0},
			"mode":         "markers+text",
			"text":         []string{label},
			"textposition": "top right",
			"textfont":     m{"color": "#ffd166", "size": 11, "family": "Fira Code, monospace"},
			"marker":       m{"color": "#ffd166", "size": 9, "line": m{"color": "#0a1628", "width": 2}},
			"name":         "Bound",
			"hoverinfo":    "text",
		})
	}

	return traces
}

func buildLayout(vp core.Viewport) m {
	axis := func(title string, rng []float64) m {
		return m{
			"range":          rng,
			"gridcolor":      "rgba(100,140,200,0.07)",
			"gridwidth":      1,
			"zerolinecolor":  "rgba(100,180,255,0.25)",
			"zerolinewidth":  1.5,
			"tickfont":       m{"color": "#4a6a9a", "size": 11},
			"title":          m{"text": title, "font": m{"color": "#6a9acc", "size": 14}},
			"showspikes":     true,
			"spikecolor":     "rgba(0,230,160,0.3)",
			"spikethickness": 1,
			"spikedash":      "dot",
			"linecolor":      "rgba(100,140,200,0.15)",
		}
	}
	return m{
		"paper_bgcolor": "#06101e",
		"plot_bgcolor":  "#06101e",
		"font":          m{"color": "#8ba8cc", "family": "'Space Mono', 'Courier New', monospace", "size": 12},
		"xaxis":         axis("x", []float64{vp.XMin, vp.XMax}),
		"yaxis":         axis("y", []float64{vp.YMin, vp.YMax}),
		"legend": m{
			"bgcolor":     "rgba(8,18,38,0.85)",
			"bordercolor": "rgba(100,150,220,0.15)",
			"borderwidth": 1,
			"font":        m{"color": "#8ba8cc", "size": 11, "family": "'Space Mono', monospace"},
			"x":           0.01, "y": 0.99,
			"xanchor": "left", "yanchor": "top",
		},
		"margin":    m{"l": 55, "r": 25, "t": 25, "b": 55},
		"hovermode": "closest",
		"dragmode":  "zoom",
		"modebar":   m{"bgcolor": "rgba(8,18,38,0)", "color": "#4a6a9a", "activecolor": "#00e6a0"},
	}
}

func buildFuncStr(a, b float64) string {
	const eps = 1e-12
	if math.Abs(a) < eps {
		return fmtV(b)
	}
	s := ""
	switch {
	case a == 1:
		s = "x"
	case a == -1:
		s = "-x"
	default:
		s = fmtV(a) + "x"
	}
	if b > eps {
		s += " + " + fmtV(b)
	} else if b < -eps {
		s += " - " + fmtV(-b)
	}
	return s
}

func fmtV(v float64) string {
	v = math.Round(v*1000) / 1000
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}
	return fmt.Sprintf("%g", v)
}
