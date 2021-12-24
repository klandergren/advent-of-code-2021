package day11

import (
	"strings"
	"testing"
)

var testInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	totalFlashes, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if totalFlashes != 1656 {
		t.Errorf("totalFlashes for test input should be 1656, got: %d", totalFlashes)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	stepNum, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if stepNum != 195 {
		t.Errorf("stepNum for test input should be 195, got: %d", stepNum)
	}

}
