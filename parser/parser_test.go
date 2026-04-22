package parser

import (
	"go-math-flow/core"
	equat "go-math-flow/topics/linear_equations"
	"testing"
)

func TestParse_LinearEquation(t *testing.T) {
	p, err := Parse("2x + 3 = 7")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Kind() != core.KindLinearEquation {
		t.Fatalf("expected KindLinearEquation, got %d", p.Kind())
	}
	lp := p.(equat.LinearProblem)
	if lp.A != 2 || lp.B != -4 {
		t.Fatalf("expected A=2 B=-4, got A=%g B=%g", lp.A, lp.B)
	}
}

func TestParse_NoOperator(t *testing.T) {
	_, err := Parse("2x + 3")
	if err == nil {
		t.Fatal("expected error for missing operator")
	}
}
