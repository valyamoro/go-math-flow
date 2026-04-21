package main

import (
	"encoding/json"
	"fmt"
	"html/template"
)

func buildPageData(eq Equation, xmin, xmax, ymin, ymax float64) (PageData, error) {
	kind, root := eq.Solve()

	ei := EqInfo{
		Original: eq.Original,
		Standard: buildStandardForm(eq.A, eq.B),
		FuncStr:  buildFuncStr(eq.A, eq.B),
		Kind:     kind,
	}
	if kind == SolutionUnique {
		ei.Root = r3(root)
		ei.HasRoot = true
		ei.StepMul = buildStepMul(eq.A, eq.B)
		ei.StepDiv = buildStepDiv(eq.A, eq.B)
	}

	type m = map[string]interface{}
	var traces []m

	// Линия y = Ax + B
	const nPts = 300
	xs := make([]float64, nPts)
	ys := make([]float64, nPts)
	step := (xmax - xmin) / float64(nPts-1)
	for i := range xs {
		x := xmin + float64(i)*step
		xs[i] = x
		ys[i] = eq.A*x + eq.B
	}
	traces = append(traces, m{
		"type": "scatter",
		"x":    xs,
		"y":    ys,
		"mode": "lines",
		"name": "y = " + ei.FuncStr,
		"line": m{"color": "#4EC9DC", "width": 2.8},
	})

	if kind == SolutionUnique {
		traces = append(traces, m{
			"type":       "scatter",
			"x":          []float64{ei.Root, ei.Root},
			"y":          []float64{ymin, ymax},
			"mode":       "lines",
			"name":       "",
			"showlegend": false,
			"line":       m{"color": "rgba(0,230,160,0.22)", "width": 1, "dash": "dot"},
			"hoverinfo":  "skip",
		})

		label := fmt.Sprintf("x = %s", fmtF(ei.Root))
		traces = append(traces, m{
			"type":         "scatter",
			"x":            []float64{ei.Root},
			"y":            []float64{0},
			"mode":         "markers+text",
			"text":         []string{label},
			"textposition": "top right",
			"textfont":     m{"color": "#00e6a0", "size": 11, "family": "Fira Code, monospace"},
			"marker":       m{"color": "#00e6a0", "size": 9, "line": m{"color": "#0a1628", "width": 2}},
			"name":         "Root",
			"hoverinfo":    "text",
		})
	}

	layout := m{
		"paper_bgcolor": "#06101e",
		"plot_bgcolor":  "#06101e",
		"font":          m{"color": "#8ba8cc", "family": "'Space Mono', 'Courier New', monospace", "size": 12},
		"xaxis": m{
			"range":          []float64{xmin, xmax},
			"gridcolor":      "rgba(100,140,200,0.07)",
			"gridwidth":      1,
			"zerolinecolor":  "rgba(100,180,255,0.25)",
			"zerolinewidth":  1.5,
			"tickfont":       m{"color": "#4a6a9a", "size": 11},
			"title":          m{"text": "x", "font": m{"color": "#6a9acc", "size": 14}},
			"showspikes":     true,
			"spikecolor":     "rgba(0,230,160,0.3)",
			"spikethickness": 1,
			"spikedash":      "dot",
			"linecolor":      "rgba(100,140,200,0.15)",
		},
		"yaxis": m{
			"range":          []float64{ymin, ymax},
			"gridcolor":      "rgba(100,140,200,0.07)",
			"gridwidth":      1,
			"zerolinecolor":  "rgba(100,180,255,0.25)",
			"zerolinewidth":  1.5,
			"tickfont":       m{"color": "#4a6a9a", "size": 11},
			"title":          m{"text": "y", "font": m{"color": "#6a9acc", "size": 14}},
			"showspikes":     true,
			"spikecolor":     "rgba(0,230,160,0.3)",
			"spikethickness": 1,
			"spikedash":      "dot",
			"linecolor":      "rgba(100,140,200,0.15)",
		},
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
		"modebar": m{
			"bgcolor":     "rgba(8,18,38,0)",
			"color":       "#4a6a9a",
			"activecolor": "#00e6a0",
		},
	}

	tj, err := json.Marshal(traces)
	if err != nil {
		return PageData{}, err
	}
	lj, err := json.Marshal(layout)
	if err != nil {
		return PageData{}, err
	}
	return PageData{
		TracesJSON: template.JS(tj),
		LayoutJSON: template.JS(lj),
		Eq:         ei,
		XMin:       xmin, XMax: xmax,
		YMin: ymin, YMax: ymax,
	}, nil
}
