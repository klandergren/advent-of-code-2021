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

func Test_NewCaveNetwork(t *testing.T) {
	r := strings.NewReader(testInput1)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		t.Error(err)
	}

	numCaves := len(cn.Caves)
	if numCaves != 6 {
		t.Errorf("cavern should have 6 caves but got: %d\n", numCaves)
	}

	numTunnels := len(cn.Tunnels)
	if numTunnels != 14 {
		t.Errorf("cavern should have 14 tunnels but got: %d\n", numTunnels)
	}

}

func Test_AllPaths1(t *testing.T) {
	r := strings.NewReader(testInput1)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		t.Error(err)
	}

	numAllPaths := len(cn.AllPaths())

	if numAllPaths != 10 {
		t.Errorf("cavern should have 10 paths but got: %d\n", numAllPaths)
	}
}

func Test_AllPaths2(t *testing.T) {
	r := strings.NewReader(testInput2)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		t.Error(err)
	}

	numAllPaths := len(cn.AllPaths())

	if numAllPaths != 19 {
		t.Errorf("cavern should have 19 paths but got: %d\n", numAllPaths)
	}
}

func Test_AllPaths3(t *testing.T) {
	r := strings.NewReader(testInput3)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		t.Error(err)
	}

	numAllPaths := len(cn.AllPaths())

	if numAllPaths != 226 {
		t.Errorf("cavern should have 226 paths but got: %d\n", numAllPaths)
	}
}
