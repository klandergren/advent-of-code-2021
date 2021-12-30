package day13

import (
	"strings"
	"testing"
)

var inputData = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(inputData)

	count, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if count != 17 {
		t.Errorf("expected dots visible after applying first instruction to be 17 but got: %d", count)
	}
}
