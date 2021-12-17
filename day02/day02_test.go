package day02

import (
	"strings"
	"testing"
)

var testInput = `forward 5
down 5
forward 8
up 3
down 8
forward 2
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	product, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if product != 150 {
		t.Errorf("product should be 150, got: %d", product)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	product, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if product != 900 {
		t.Errorf("product should be 900, got: %d", product)
	}

}
