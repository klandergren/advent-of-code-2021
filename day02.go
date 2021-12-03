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

type Submarine struct {
	Aim, X, Depth int
}

func (sub *Submarine) Mult() int {
	return sub.X * sub.Depth
}

func (sub *Submarine) AdvanceWithoutAim(step string) {
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
		sub.X = sub.X + magnitude
	case "down":
		sub.Depth = sub.Depth + magnitude
	case "up":
		sub.Depth = sub.Depth - magnitude
	default:
		log.Fatal("how'd we get here?")
	}

}

func (sub *Submarine) AdvanceWithAim(step string) {
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
		sub.X = sub.X + magnitude
		sub.Depth = sub.Depth + sub.Aim*magnitude
	case "down":
		sub.Aim = sub.Aim + magnitude
	case "up":
		sub.Aim = sub.Aim - magnitude
	default:
		log.Fatal("how'd we get here?")
	}

	//	fmt.Println("step: ", step)
	//	fmt.Println("sub: %v", sub)

}

func part1() {
	//	file, err := os.Open("./test-data/day02.txt")
	file, err := os.Open("./input-data/day02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	sub := Submarine{0, 0, 0}

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		sub.AdvanceWithoutAim(string(line))
	}

	fmt.Println("part 1: ", sub.Mult())
}

func part2() {
	// file, err := os.Open("./test-data/day02.txt")
	file, err := os.Open("./input-data/day02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	sub := Submarine{0, 0, 0}

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		sub.AdvanceWithAim(string(line))
	}

	fmt.Println("part 2: ", sub.Mult())
}
