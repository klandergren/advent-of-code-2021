package day01

import (
	"strings"
	"testing"
)

var testInput = `199
200
208
210
200
207
240
269
260
263
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	count, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if count != 7 {
		t.Errorf("count should be 7, got: %d", count)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	count, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if count != 5 {
		t.Errorf("count should be 5, got: %d", count)
	}

}
