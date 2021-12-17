package day03

import (
	"strings"
	"testing"
)

var testInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	consumption, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if consumption != 198 {
		t.Errorf("consumption should be 198, got: %d", consumption)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	lsr, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if lsr != 230 {
		t.Errorf("life support rating should be 230, got: %d", lsr)
	}

}
