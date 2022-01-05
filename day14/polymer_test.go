package day14

import "testing"

func Test_Polymer(t *testing.T) {
	template := []rune{'A', 'B', 'C', 'D'}

	rules := []Rule{
		NewRule("AB -> E"),
		NewRule("BC -> F"),
		NewRule("CD -> G"),
	}

	p := NewPolymer(template, rules)
	p.ApplySteps(1)

	if string(p.Sequence) != "AEBFCGD" {
		t.Fatalf("expected AEBFCGD but got: %v\n", string(p.Sequence))
	}

}
