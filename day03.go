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

func part1() {
	//file, err := os.Open("./test-data/day03.txt")
	file, err := os.Open("./input-data/day03.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	grid := make([][]int, 0)

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		for x, rune := range string(line) {
			bit, err := strconv.Atoi(string(rune))

			if err != nil {
				log.Fatal(err)
			}

			if len(grid) <= x {
				grid = append(grid, make([]int, 0))
			}

			grid[x] = append(grid[x], bit)
		}
	}

	gammaRaw := make([]string, len(grid))

	for x, col := range grid {

		onesCount := 0
		zerosCount := 0

		for _, y := range col {
			if y == 0 {
				zerosCount++
			} else {
				onesCount++
			}
		}

		if onesCount < zerosCount {
			gammaRaw[x] = "0"
		} else if zerosCount < onesCount {
			gammaRaw[x] = "1"
		} else {
			log.Fatal("equal?")
		}
	}

	gamma, err := strconv.ParseInt(strings.Join(gammaRaw, ""), 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	epsilonRaw := make([]string, len(grid))

	for i, b := range gammaRaw {
		if b == "0" {
			epsilonRaw[i] = "1"
		} else {
			epsilonRaw[i] = "0"
		}
	}

	epsilon, err := strconv.ParseInt(strings.Join(epsilonRaw, ""), 2, 64)

	consumption := gamma * epsilon

	fmt.Println("part 1: ", consumption)
}

func part2() {
}
