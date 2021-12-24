package day12

import (
	"strings"
	"testing"

	"github.com/klandergren/advent-of-code-2021/util"
)

var testInput1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

var testInput2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

var testInput3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

func Test_PartOne(t *testing.T) {
	AllPaths(t, testInput1, false, 10)
	AllPaths(t, testInput2, false, 19)
	AllPaths(t, testInput3, false, 226)
}

func Test_PartTwo(t *testing.T) {
	AllPaths(t, testInput1, true, 36)
	AllPaths(t, testInput2, true, 103)
	AllPaths(t, testInput3, true, 3509)
}

func AllPaths(t *testing.T, input string, shouldAllowDouble bool, expected int) {
	r := strings.NewReader(input)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		t.Error(err)
	}

	numAllPaths := len(cn.AllPaths(shouldAllowDouble))

	if numAllPaths != expected {
		t.Errorf("cavern should have %d paths but got: %d\n", expected, numAllPaths)
	}
}
