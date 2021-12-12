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

func MinMax(s []int) (min int, max int) {
	if s == nil {
		log.Fatal("nil slice")
	}

	if len(s) == 0 {
		log.Fatal("empty slice")
	}

	min = s[0]
	max = s[0]

	for i := 1; i < len(s); i++ {
		maybeMin := s[i]
		maybeMax := s[i]

		if maybeMin < min {
			min = maybeMin
		}

		if max < maybeMax {
			max = maybeMax
		}

	}

	return min, max
}

func FuelCost(positions []int, alignmentPosition int) int {
	if positions == nil {
		log.Fatal("nil slice")
	}

	if len(positions) == 0 {
		log.Fatal("empty slice")
	}

	min, max := MinMax(positions)

	if alignmentPosition < min {
		log.Fatal("too low")
	}

	if max < alignmentPosition {
		log.Fatal("too high")
	}

	cost := 0

	for _, p := range positions {
		if p <= alignmentPosition {
			cost = cost + (alignmentPosition - p)
		} else {
			cost = cost + (p - alignmentPosition)
		}
	}

	return cost
}

func FindMinFuelCostPositionBrute(positions []int) int {
	min, max := MinMax(positions)

	minFuelCost := FuelCost(positions, min)

	for i := min + 1; i <= max; i++ {
		maybeMinFuelCost := FuelCost(positions, i)
		if maybeMinFuelCost < minFuelCost {
			minFuelCost = maybeMinFuelCost
		}
	}

	return minFuelCost
}

func part1() {
	fmt.Println("part 1: ", FindMinFuelCostPositionBrute(load()))
}

func part2() {
	fmt.Println("part 2: ", 0)
}

func load() (horizontalPositions []int) {
	//file, err := os.Open("./test-data/day07.txt")
	file, err := os.Open("./input-data/day07.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		for _, sPos := range strings.Split(string(line), ",") {
			pos, err := strconv.Atoi(sPos)

			if err != nil {
				log.Fatal(err)
			}

			horizontalPositions = append(horizontalPositions, pos)
		}

	}

	return horizontalPositions
}
