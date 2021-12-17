package day05

import (
	"io"
	"strconv"
	"strings"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	ventLines := createVentLines(lines)
	grid := NewGridWithoutDiagonals(ventLines)

	return grid.NumOverlappingAtLeast(2), nil
}

func PartTwo(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	ventLines := createVentLines(lines)

	grid := NewGridWithDiagonals(ventLines)

	return grid.NumOverlappingAtLeast(2), nil
}

func createVentLines(lines []string) (ventLines []*VentLine) {
	for _, line := range lines {
		coords := strings.Split(line, " -> ")

		coord0 := strings.Split(coords[0], ",")
		coord1 := strings.Split(coords[1], ",")

		x1, _ := strconv.Atoi(coord0[0])
		y1, _ := strconv.Atoi(coord0[1])
		x2, _ := strconv.Atoi(coord1[0])
		y2, _ := strconv.Atoi(coord1[1])

		ventLines = append(ventLines, &VentLine{
			&Coord{x1, y1},
			&Coord{x2, y2},
		})
	}
	return ventLines
}
