package day06

import (
	"strings"
	"testing"
)

var testInput = `3,4,3,1,2
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	count, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if count != 5934 {
		t.Errorf("count should be 5934, got: %d", count)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	count, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if count != 26984457539 {
		t.Errorf("count should be 26984457539, got: %d", count)
	}

}
