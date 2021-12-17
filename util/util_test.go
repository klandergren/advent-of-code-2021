package util

import (
	"strings"
	"testing"
)

var testInput = `line 1
line 2
line 3
`

func Test_LoadLines(t *testing.T) {
	r := strings.NewReader(testInput)

	lines, err := LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	count := len(lines)
	if count != 3 {
		t.Errorf("count should be 3, got: %d", count)
	}

}
