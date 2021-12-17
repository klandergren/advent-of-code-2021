package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartOne(reader *bufio.Reader) (int, error) {
	ventLines := load(reader)

	grid := NewGridWithoutDiagonals(ventLines)

	return grid.NumOverlappingAtLeast(2), nil
}
func PartTwo(reader *bufio.Reader) (int, error) {
	ventLines := load(reader)

	grid := NewGridWithDiagonals(ventLines)

	return grid.NumOverlappingAtLeast(2), nil
}

func load(reader *bufio.Reader) (ventLines []*VentLine) {
	ventLines = make([]*VentLine, 0)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		coords := strings.Split(string(line), " -> ")

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
