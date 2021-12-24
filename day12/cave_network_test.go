package day12

import (
	"strings"
	"testing"

	"github.com/klandergren/advent-of-code-2021/util"
)

var testInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

func Test_NewCaveNetwork(t *testing.T) {
	r := strings.NewReader(testInput)
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
