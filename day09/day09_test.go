package day09

import (
	"strings"
	"testing"
)

func Test_SumRisk(t *testing.T) {
	p1 := &Point{0, 0, 0}
	p2 := &Point{0, 0, 0}

	if sum := SumRisk([]*Point{p1, p2}); sum != 2 {
		t.Errorf("sum risk is %d; want 2", sum)
	}
}

func Test_load(t *testing.T) {
	hm, err := load(strings.NewReader("2199943210\n3987894921\n"))
	t.Log(hm)
	t.Log(err)

}
