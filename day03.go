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
	//	file, err := os.Open("./test-data/day03.txt")
	file, err := os.Open("./input-data/day03.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	gridXY := make([][]int, 0)
	gridYX := make([][]int, 0)

	y := 0
	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		if len(gridYX) <= y {
			gridYX = append(gridYX, make([]int, len(string(line))))
		}

		for x, rune := range string(line) {
			bit, err := strconv.Atoi(string(rune))

			if err != nil {
				log.Fatal(err)
			}

			if len(gridXY) <= x {
				gridXY = append(gridXY, make([]int, 0))
			}

			gridXY[x] = append(gridXY[x], bit)
			gridYX[y][x] = bit
		}
		y++
	}

	oxy := findOxygen(gridYX)
	co2 := findCO2(gridYX)
	fmt.Println("oxy: ", oxy)
	fmt.Println("co2: ", co2)

	fmt.Println("part 2: ", oxy*co2)
}

func findOxygen(gridYX [][]int) int {
	// filter
	oxy := filterOxy(gridYX, 0)

	oxyRaw := make([]string, len(oxy))
	for i, v := range oxy {
		oxyRaw[i] = strconv.Itoa(v)
	}

	fmt.Println(oxyRaw)
	// convert
	ox, _ := strconv.ParseInt(strings.Join(oxyRaw, ""), 2, 64)
	return int(ox)
}

func filterOxy(candidates [][]int, pos int) []int {
	fmt.Println("candidates:", candidates)
	fmt.Println("pos:", pos)

	if len(candidates) == 1 {
		return candidates[0]
	}

	onesCount := 0
	zerosCount := 0

	for _, col := range candidates {
		if col[pos] == 0 {
			zerosCount++
		} else {
			onesCount++
		}
	}

	mostCommon := 0
	if onesCount < zerosCount {
		mostCommon = 0
	} else if zerosCount < onesCount {
		mostCommon = 1
	} else {
		mostCommon = 1
	}

	indexesToKeep := make([]int, 0)
	for i, c := range candidates {
		if c[pos] == mostCommon {
			indexesToKeep = append(indexesToKeep, i)
		}
	}
	fmt.Println("indexesToKeep:", indexesToKeep)

	filteredCandidates := make([][]int, len(indexesToKeep))

	for i, j := range indexesToKeep {
		filteredCandidates[i] = candidates[j]
	}
	fmt.Println("filteredCandidates:", filteredCandidates)

	return filterOxy(filteredCandidates, pos+1)
}

func findCO2(gridYX [][]int) int {
	// filter
	co2 := filterCO2(gridYX, 0)

	co2Raw := make([]string, len(co2))
	for i, v := range co2 {
		co2Raw[i] = strconv.Itoa(v)
	}

	fmt.Println(co2Raw)
	// convert
	co, _ := strconv.ParseInt(strings.Join(co2Raw, ""), 2, 64)
	return int(co)
}

func filterCO2(candidates [][]int, pos int) []int {
	fmt.Println("candidates:", candidates)
	fmt.Println("pos:", pos)

	if len(candidates) == 1 {
		return candidates[0]
	}

	onesCount := 0
	zerosCount := 0

	for _, col := range candidates {
		if col[pos] == 0 {
			zerosCount++
		} else {
			onesCount++
		}
	}

	leastCommon := 0
	if onesCount < zerosCount {
		leastCommon = 1
	} else if zerosCount < onesCount {
		leastCommon = 0
	} else {
		leastCommon = 0
	}
	fmt.Println("leastCommon: ", leastCommon)

	indexesToKeep := make([]int, 0)
	for i, c := range candidates {
		if c[pos] == leastCommon {
			indexesToKeep = append(indexesToKeep, i)
		}
	}
	fmt.Println("indexesToKeep:", indexesToKeep)

	filteredCandidates := make([][]int, len(indexesToKeep))

	for i, j := range indexesToKeep {
		filteredCandidates[i] = candidates[j]
	}
	fmt.Println("filteredCandidates:", filteredCandidates)

	return filterCO2(filteredCandidates, pos+1)
}
