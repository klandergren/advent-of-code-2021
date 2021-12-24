package day12

import (
	"strings"
)

type CaveNetwork struct {
	Caves   map[string]Cave
	Tunnels []*Tunnel
}

func NewCaveNetwork(lines []string) (*CaveNetwork, error) {
	caves := make(map[string]Cave)
	tunnels := make([]*Tunnel, 0)

	// each line is a cave / tunnel
	for _, line := range lines {
		names := strings.Split(line, "-")

		for _, name := range names {
			// do we already have it?
			if _, ok := caves[name]; !ok {
				caves[name] = NewCave(name)
			}
		}

		caveA := caves[names[0]]
		caveB := caves[names[1]]

		tunnels = append(tunnels, &Tunnel{caveA, caveB})
		tunnels = append(tunnels, &Tunnel{caveB, caveA})
	}

	cn := &CaveNetwork{caves, tunnels}

	return cn, nil
}

func (cn *CaveNetwork) AllPaths(shouldAllowDouble bool) [][]string {
	cw := &CaveWalker{}
	paths := cw.WalkAll(cn, "start", "end", shouldAllowDouble)

	// for _, p := range paths {
	// 	fmt.Println(strings.Join(p, ","))
	// }
	return paths
}
