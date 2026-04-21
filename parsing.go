package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ParseEquation парсит уравнение вида "2x + 3 = 7" или "-x + 4 = 2x - 1"
// и приводит к стандартной форме Ax + B = 0.
func ParseEquation(s string) (Equation, error) {
	s = strings.TrimSpace(s)
	idx := strings.Index(s, "=")
	if idx < 0 {
		return Equation{}, fmt.Errorf("no '=' in %q", s)
	}
	lhsRaw := strings.TrimSpace(s[:idx])
	rhsRaw := strings.TrimSpace(s[idx+1:])

	la, lb, err := parseLinearExpr(lhsRaw)
	if err != nil {
		return Equation{}, fmt.Errorf("LHS %q: %w", lhsRaw, err)
	}
	ra, rb, err := parseLinearExpr(rhsRaw)
	if err != nil {
		return Equation{}, fmt.Errorf("RHS %q: %w", rhsRaw, err)
	}

	// (la)x + lb = (ra)x + rb  →  (la-ra)x + (lb-rb) = 0
	return Equation{A: la - ra, B: lb - rb, Original: s}, nil
}

// parseLinearExpr разбирает выражение вида "2x + 3" или "-x - 5"
// и возвращает (коэффициент при x, свободный член).
func parseLinearExpr(expr string) (coeffX, constant float64, err error) {
	expr = strings.ReplaceAll(expr, " ", "")
	if expr == "" {
		return 0, 0, nil
	}
	if expr[0] != '+' && expr[0] != '-' {
		expr = "+" + expr
	}
	re := regexp.MustCompile(`[+-][^+-]+`)
	for _, tok := range re.FindAllString(expr, -1) {
		if strings.ContainsRune(tok, 'x') {
			coeffX += extractCoeff(strings.ReplaceAll(tok, "x", ""))
		} else {
			v, e := strconv.ParseFloat(tok, 64)
			if e != nil {
				err = fmt.Errorf("unrecognised token %q", tok)
				return
			}
			constant += v
		}
	}
	return
}

func extractCoeff(s string) float64 {
	switch s {
	case "+", "":
		return 1
	case "-":
		return -1
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
