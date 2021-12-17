package day09

import "testing"

func Test_Risk(t *testing.T) {
	p := &Point{0, 0, 0}

	if r := p.Risk(); r != 1 {
		t.Errorf("%s risk is %d; want 1", p, r)
	}
}
