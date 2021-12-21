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
	t.SkipNow()
	r := strings.NewReader(testInput)

	score, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if score != 26397 {
		t.Errorf("score for test input should be 26397, got: %d", score)
	}

}

func Test_PartTwo(t *testing.T) {
	t.SkipNow()
	r := strings.NewReader(testInput)

	score, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if score != 288957 {
		t.Errorf("score for test input should be 288957, got: %d", score)
	}

}
