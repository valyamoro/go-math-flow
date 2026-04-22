package parser

import (
	"go-math-flow/core"
	linear "go-math-flow/topics/linear_equations"
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
	lp := p.(linear.LinearProblem)
	if lp.A != 2 || lp.B != -4 {
		t.Fatalf("expected A=2 B=-4, got A=%g B=%g", lp.A, lp.B)
	}
}

func TestParse_LinearInequality_GT(t *testing.T) {
	p, err := Parse("3x - 6 > 0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Kind() != core.KindLinearInequality {
		t.Fatalf("expected KindLinearInequality, got %d", p.Kind())
	}
	lp := p.(linear.LinearProblem)
	if lp.Op != ">" {
		t.Fatalf("expected op '>', got %q", lp.Op)
	}
}

func TestParse_LinearInequality_LTE(t *testing.T) {
	p, err := Parse("x + 1 <= 5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lp := p.(linear.LinearProblem)
	if lp.Op != "<=" || lp.A != 1 || lp.B != -4 {
		t.Fatalf("got op=%q A=%g B=%g", lp.Op, lp.A, lp.B)
	}
}

func TestParse_NoOperator(t *testing.T) {
	_, err := Parse("2x + 3")
	if err == nil {
		t.Fatal("expected error for missing operator")
	}
}
