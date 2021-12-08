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

type LanternFish struct {
	SpawnTimer int
}

func NewLanternFish(st int) (lf *LanternFish) {
	return &LanternFish{
		SpawnTimer: st,
	}
}

func (lf *LanternFish) Advance() (child *LanternFish) {
	if lf.SpawnTimer == 0 {
		child = NewLanternFish(8)
		lf.SpawnTimer = 6
	} else {
		lf.SpawnTimer--
	}
	return child
}

func (lf *LanternFish) String() string {
	return fmt.Sprintf("%d", lf.SpawnTimer)
}

func part1() {
	school := load()

	maxDays := 80
	day := 1

	//	fmt.Println("initial state: ", school)
	for day <= maxDays {

		nextGeneration := make([]*LanternFish, 0)
		for _, lf := range school {
			child := lf.Advance()

			if child != nil {
				nextGeneration = append(nextGeneration, child)
			}

		}

		if 0 < len(nextGeneration) {
			school = append(school, nextGeneration...)
		}

		//		fmt.Println("after ", day, ": ", school)
		day++
	}

	fmt.Println("part 1: ", len(school))
}

func part2() {
	fmt.Println("part 2: ", 0)
}

func load() (school []*LanternFish) {
	//file, err := os.Open("./test-data/day06.txt")
	file, err := os.Open("./input-data/day06.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	school = make([]*LanternFish, 0)
	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		spawnTimers := strings.Split(string(line), ",")

		for _, spawnTimer := range spawnTimers {
			st, err := strconv.Atoi(spawnTimer)

			if err != nil {
				log.Fatal(err)
			}

			school = append(school, NewLanternFish(st))
		}

	}

	return school
}
