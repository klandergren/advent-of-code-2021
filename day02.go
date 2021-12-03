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
	parta()
	partb()

	fmt.Println("done")
}

type Coordinate struct {
	X, Depth int
}

func (c *Coordinate) Mult() int {
	return c.X * c.Depth
}

func (c *Coordinate) Advance(step string) {
	parts := strings.Fields(step)

	if len(parts) != 2 {
		log.Fatal("unexpected step format: ", step)
	}

	direction := parts[0]
	magnitude, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "forward":
		c.X = c.X + magnitude
	case "down":
		c.Depth = c.Depth + magnitude
	case "up":
		c.Depth = c.Depth - magnitude
	default:
		log.Fatal("how'd we get here?")
	}

}

func parta() {
	//	file, err := os.Open("./test-data/day02.txt")
	file, err := os.Open("./input-data/day02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	coord := Coordinate{0, 0}

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		coord.Advance(string(line))
	}

	fmt.Println("part a: ", coord.Mult())
}

func partb() {
}
