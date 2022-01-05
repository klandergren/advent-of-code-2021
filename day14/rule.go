package day14

import "strings"

type Rule struct {
	A, B, C rune
}

func NewRule(input string) Rule {
	parts := strings.Split(input, " -> ")

	matching := []rune(parts[0])
	a := matching[0]
	b := matching[1]

	element := []rune(parts[1])
	c := element[0]

	return Rule{
		A: a,
		B: b,
		C: c,
	}
}

func (r Rule) Matches(pair []rune) bool {
	return (pair[0] == r.A) && (pair[1] == r.B)
}

func (r Rule) Element() rune {
	return r.C
}
