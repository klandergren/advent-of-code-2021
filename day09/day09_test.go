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
