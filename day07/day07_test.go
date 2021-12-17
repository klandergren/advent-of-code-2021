package day07

import (
	"strings"
	"testing"
)

var testInput = `16,1,2,0,4,2,7,1,2,14
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	cost, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if cost != 37 {
		t.Errorf("cost should be 37, got: %d", cost)
	}

}

func Test_PartTwo(t *testing.T) {
	r := strings.NewReader(testInput)

	cost, err := PartTwo(r)
	if err != nil {
		t.Error(err)
	}

	if cost != 168 {
		t.Errorf("cost should be 168, got: %d", cost)
	}

}
