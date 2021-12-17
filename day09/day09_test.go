package day09

import (
	"strings"
	"testing"
)

var testInput = `2199943210
3987894921
9856789892
8767896789
9899965678
`

func Test_PartOne(t *testing.T) {
	reader := strings.NewReader(testInput)

	sumRisk, err := PartOne(reader)

	if err != nil {
		t.Error(err)
	}

	if sumRisk != 15 {
		t.Errorf("sum risk for test height map should be 15, got: %d", sumRisk)
	}

}

func Test_PartTwo(t *testing.T) {
	reader := strings.NewReader(testInput)

	product, err := PartTwo(reader)

	if err != nil {
		t.Error(err)
	}

	if product != 1134 {
		t.Errorf("product for basin sizes should be 1134, got: %d", product)
	}

}
