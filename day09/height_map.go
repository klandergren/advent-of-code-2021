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

func (hm HeightMap) LowPointHeights() (pts []Height) {
	for y, row := range hm {
		for x, h := range row {
			isLowPoint := true
			for _, ah := range hm.AdjacentHeights(x, y) {
				isLowPoint = isLowPoint && (h < ah)
			}
			if isLowPoint {
				pts = append(pts, h)
			}
		}
	}

	return pts
}

func (hm HeightMap) AdjacentHeights(x, y int) (adjacentPoints []Height) {
	c := &Coordinate{x, y}
	for _, ac := range c.AdjacentCoordinates() {
		if h := hm.Get(ac.X, ac.Y); h != nil {
			adjacentPoints = append(adjacentPoints, *h)
		}
	}
	return adjacentPoints
}

func (hm HeightMap) Get(x, y int) *Height {
	if y < 0 {
		return nil
	}

	if len(hm) <= y {
		return nil
	}

	if x < 0 {
		return nil
	}

	if len(hm[y]) <= x {
		return nil
	}

	return &hm[y][x]
}
