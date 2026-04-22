package parser

import (
	"fmt"
	"go-math-flow/core"
	linear "go-math-flow/topics/linear_equations"
	ineq "go-math-flow/topics/linear_inequalities"
	"regexp"
	"strconv"
	"strings"
)

var opRe = regexp.MustCompile(`<=|>=|<|>|=`)

func Parse(s string) (core.MathProblem, error) {
	s = strings.TrimSpace(s)

	loc := opRe.FindStringIndex(s)
	if loc == nil {
		return nil, fmt.Errorf("no relation operator found in %q", s)
	}
	op := s[loc[0]:loc[1]]
	lhsRaw := strings.TrimSpace(s[:loc[0]])
	rhsRaw := strings.TrimSpace(s[loc[1]:])

	lax, lay, lb, err := parseLinearExpr(lhsRaw)
	if err != nil {
		return nil, fmt.Errorf("left side %q: %w", lhsRaw, err)
	}
	rax, ray, rb, err := parseLinearExpr(rhsRaw)
	if err != nil {
		return nil, fmt.Errorf("right side %q: %w", rhsRaw, err)
	}

	if op == "=" {
		return linear.New(lax-rax, lay-ray, lb-rb, op, s), nil
	}

	return ineq.New(lax-rax, lay-ray, lb-rb, op, s), nil
}

func parseLinearExpr(expr string) (coeffX, coeffY, constant float64, err error) {
	expr = strings.ReplaceAll(expr, " ", "")
	if expr == "" {
		return 0, 0, 0, nil
	}
	if expr[0] != '+' && expr[0] != '-' {
		expr = "+" + expr
	}
	re := regexp.MustCompile(`[+-][^+-]+`)
	for _, tok := range re.FindAllString(expr, -1) {
		switch {
		case strings.ContainsRune(tok, 'x'):
			coeffX += extractCoeff(strings.ReplaceAll(tok, "x", ""))
		case strings.ContainsRune(tok, 'y'):
			coeffY += extractCoeff(strings.ReplaceAll(tok, "y", ""))
		default:
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
