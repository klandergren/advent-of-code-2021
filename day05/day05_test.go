package day05

import (
	"strings"
	"testing"
)

var testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	count, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if count != 5 {
		t.Errorf("count should be 5, got: %d", count)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	count, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if count != 12 {
		t.Errorf("count should be 12, got: %d", count)
	}

}
