package day05

import (
	"fmt"
	"strings"
)

type Grid struct {
	MaxX, MaxY int
	ReadingsYX [][]*Reading
}

func NewGridWithoutDiagonals(ventLines []*VentLine) *Grid {
	return NewGrid(ventLines, true)
}

func NewGridWithDiagonals(ventLines []*VentLine) *Grid {
	return NewGrid(ventLines, false)
}

func NewGrid(ventLines []*VentLine, skipDiagonals bool) *Grid {
	maxX, maxY := 0, 0

	for _, vl := range ventLines {
		if skipDiagonals && vl.IsDiagonal() {
			continue
		}

		if maxX < vl.MaxX() {
			maxX = vl.MaxX()
		}

		if maxY < vl.MaxY() {
			maxY = vl.MaxY()
		}

	}

	// populate initial board
	readingsYX := make([][]*Reading, maxY+1)
	for y, _ := range readingsYX {
		readingsYX[y] = make([]*Reading, maxX+1)

		for x := range readingsYX[y] {
			readingsYX[y][x] = &Reading{
				0,
			}
		}
	}

	// populate gridlines
	for _, vl := range ventLines {
		if skipDiagonals && vl.IsDiagonal() {
			continue
		}

		for _, coord := range vl.Coords() {
			reading := readingsYX[coord.Y][coord.X]
			reading.Inc()
		}
	}

	grid := &Grid{maxX, maxY, readingsYX}

	fmt.Println(grid)

	return grid
}

func (g *Grid) NumOverlappingAtLeast(min int) int {
	count := 0

	for y, row := range g.ReadingsYX {
		for x := range row {
			reading := g.ReadingsYX[y][x]
			if min <= reading.NumVents {
				count++
			}
		}
	}

	return count
}

func (g *Grid) String() string {
	var sb strings.Builder

	sb.WriteString("\n")

	for y := 0; y < len(g.ReadingsYX); y++ {
		for x := 0; x < len(g.ReadingsYX[y]); x++ {
			reading := g.ReadingsYX[y][x]
			sb.WriteString(reading.String())
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
