package day12

import (
	"fmt"
)

type CaveWalker struct {
}

func (cw *CaveWalker) WalkAllHelper1(cn *CaveNetwork, endName string, availableCaves []Cave, potentialPath []string, visited map[string]bool) (paths [][]string) {
	fmt.Println("called with:", endName, availableCaves, potentialPath, visited)

	// we can move
	for _, cave := range availableCaves {
		potentialPath2 := make([]string, len(potentialPath))
		copy(potentialPath2, potentialPath)
		potentialPath2 = append(potentialPath2, cave.Name)

		// are we at the end?
		if cave.Name == endName {
			paths = append(paths, potentialPath2)
			fmt.Println("at end")
			fmt.Println(paths)
			continue
		}

		// we are not at end; visit the cave
		visited2 := make(map[string]bool)
		visited2[cave.Name] = true
		for k, v := range visited {
			visited2[k] = v
		}

		// can we move anywhere?
		availableCaves2 := cw.Available(cn, cave.Name, visited2)
		fmt.Println("availableCaves2:", availableCaves2)

		// we cannot
		if len(availableCaves2) == 0 {
			fmt.Println("continuing!")
			continue
		}

		// we can move
		newPaths := cw.WalkAllHelper1(cn, endName, availableCaves2, potentialPath2, visited2)

		paths = append(paths, newPaths...)
		fmt.Println("at loop")
		fmt.Println(paths)
	}
	return paths
}

func (cw *CaveWalker) WalkAll(cn *CaveNetwork, startName, endName string) (paths [][]string) {
	potentialPath := []string{startName}
	// are we already at the end?
	if startName == endName {
		paths = append(paths, potentialPath)
		return paths
	}

	// mark our starting point
	visited := make(map[string]bool)
	visited[startName] = true

	// can we move anywhere?
	availableCaves := cw.Available(cn, startName, visited)

	// we cannot
	if len(availableCaves) == 0 {
		return paths
	}

	return cw.WalkAllHelper1(cn, endName, availableCaves, potentialPath, visited)
}

func (cw *CaveWalker) Available(cn *CaveNetwork, fromCaveName string, visited map[string]bool) (caves []Cave) {
	fmt.Println("visited:", visited)
	// find adjacent
	adj := make([]Cave, 0)
	for _, t := range cn.Tunnels {
		if fromCaveName == t.Start.Name {
			adj = append(adj, t.End)
		}
	}

	fmt.Println("adjacent to:", fromCaveName, "is:", adj)

	// filter by what we have seen
	for _, c := range adj {
		// always can revisit big caves
		if c.IsBigCave {
			caves = append(caves, c)
			continue
		}

		// only visit small caves if we haven't visited before
		if _, ok := visited[c.Name]; !ok {
			caves = append(caves, c)
		}
	}

	fmt.Println("caves is:", caves)
	return caves
}
