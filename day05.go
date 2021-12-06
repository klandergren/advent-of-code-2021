package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()

	fmt.Println("done")
}

type Reading struct {
	NumVents int
}

func (r *Reading) Inc() {
	r.NumVents++
}

// assumes 4 digit max
func (r *Reading) String() string {
	switch {
	case r.NumVents == 0:
		return "   . "
	default:
		return fmt.Sprintf("%4d ", r.NumVents)
	}
}

type Coord struct {
	X, Y int
}

func (c *Coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type VentLine struct {
	Start, End *Coord
}

func (vl *VentLine) IsDiagonal() bool {
	return vl.Start.X != vl.End.X && vl.Start.Y != vl.End.Y
}

func (vl *VentLine) MaxX() int {
	if vl.Start.X < vl.End.X {
		return vl.End.X
	} else {
		return vl.Start.X
	}
}

func (vl *VentLine) MaxY() int {
	if vl.Start.Y < vl.End.Y {
		return vl.End.Y
	} else {
		return vl.Start.Y
	}
}

// assume lines only cross whole int coordinates
func (vl *VentLine) Coords() []*Coord {
	coords := make([]*Coord, 0)

	switch {
	case vl.Start.X <= vl.End.X && vl.Start.Y <= vl.End.Y:
		for x := vl.Start.X; x <= vl.End.X; x++ {
			for y := vl.Start.Y; y <= vl.End.Y; y++ {
				coords = append(coords, &Coord{x, y})
			}
		}
	case vl.Start.X <= vl.End.X && vl.End.Y <= vl.Start.Y:
		for x := vl.Start.X; x <= vl.End.X; x++ {
			for y := vl.Start.Y; vl.End.Y <= y; y-- {
				coords = append(coords, &Coord{x, y})
			}
		}
	case vl.End.X <= vl.Start.X && vl.Start.Y <= vl.End.Y:
		for x := vl.Start.X; vl.End.X <= x; x-- {
			for y := vl.Start.Y; y <= vl.End.Y; y++ {
				coords = append(coords, &Coord{x, y})
			}
		}
	case vl.End.X <= vl.Start.X && vl.End.Y <= vl.Start.Y:
		for x := vl.Start.X; vl.End.X <= x; x-- {
			for y := vl.Start.Y; vl.End.Y <= y; y-- {
				coords = append(coords, &Coord{x, y})
			}
		}
	}
	return coords
}

func (vl *VentLine) String() string {
	return fmt.Sprintf("%s -> %s", vl.Start, vl.End)
}

type Grid struct {
	MaxX, MaxY int
	ReadingsYX [][]*Reading
}

func NewGridWithoutDiagonals(ventLines []*VentLine) *Grid {
	maxX, maxY := 0, 0

	for _, vl := range ventLines {
		// skip diagonal lines for now
		if vl.IsDiagonal() {
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
		// skip diagonal lines for now
		if vl.IsDiagonal() {
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

func part1() {
	ventLines := load()

	grid := NewGridWithoutDiagonals(ventLines)

	fmt.Println("part 1: ", grid.NumOverlappingAtLeast(2))
}

func part2() {
	fmt.Println("part 2: ", 0)
}

func load() (ventLines []*VentLine) {
	//file, err := os.Open("./test-data/day05.txt")
	file, err := os.Open("./input-data/day05.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	ventLines = make([]*VentLine, 0)
	for {
		line, _, err := r.ReadLine()

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
