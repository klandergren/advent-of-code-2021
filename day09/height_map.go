package day09

import (
	"strconv"
	"strings"
)

type HeightMap [][]Height

func NewHeightMap(lines []string) (hm HeightMap, err error) {
	for y, line := range lines {
		rowData := strings.Split(line, "")
		hm = append(hm, make([]Height, len(rowData)))
		for x, sHeight := range rowData {
			nHeight, err := strconv.Atoi(sHeight)
			if err != nil {
				return nil, err
			}

			hm[y][x] = Height(nHeight)
		}
	}

	return hm, nil
}

func (hm HeightMap) SumRisk() (sumRisk int) {
	for _, h := range hm.LowPointHeights() {
		sumRisk += h.Risk()
	}
	return sumRisk
}

func (hm HeightMap) BasinSizes() (basinSizes []int) {
	for _, c := range hm.LowPointCoordinates() {
		visited := make(map[string]bool)
		size := hm.WalkAllAdjacent(c, visited)

		basinSizes = append(basinSizes, size)
	}

	return basinSizes
}

func (hm HeightMap) WalkAllAdjacent(c Coordinate, visited map[string]bool) int {
	if _, ok := visited[c.String()]; ok {
		return 0
	}

	if h := hm[c.Y][c.X]; h == 9 {
		return 0
	}

	sum := 1
	visited[c.String()] = true

	for _, ac := range hm.AdjacentCoordinates(c.X, c.Y) {
		sum += hm.WalkAllAdjacent(ac, visited)
	}

	return sum
}

func (hm HeightMap) LowPointHeights() (lphts []Height) {
	for _, c := range hm.LowPointCoordinates() {
		lphts = append(lphts, hm[c.Y][c.X])
	}
	return lphts
}

func (hm HeightMap) LowPointCoordinates() (lpc []Coordinate) {
	for y, row := range hm {
		for x, h := range row {
			isLowPoint := true
			for _, ah := range hm.AdjacentHeights(x, y) {
				isLowPoint = isLowPoint && (h < ah)
			}
			if isLowPoint {
				lpc = append(lpc, Coordinate{x, y})
			}
		}
	}
	return lpc
}

func (hm HeightMap) AdjacentHeights(x, y int) (adjacentHeights []Height) {
	for _, ac := range hm.AdjacentCoordinates(x, y) {
		adjacentHeights = append(adjacentHeights, hm[ac.Y][ac.X])
	}
	return adjacentHeights
}

func (hm HeightMap) AdjacentCoordinates(x, y int) (adjacentCoordinates []Coordinate) {
	c := Coordinate{x, y}
	for _, ac := range c.AdjacentCoordinates() {
		if ac.Y < 0 {
			continue
		}

		if len(hm) <= ac.Y {
			continue
		}

		if ac.X < 0 {
			continue
		}

		if len(hm[ac.Y]) <= ac.X {
			continue
		}

		adjacentCoordinates = append(adjacentCoordinates, *ac)
	}
	return adjacentCoordinates
}
