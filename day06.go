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

func (lf *LanternFish) AllDescendantsAfter(day int, m map[string]int) int {
	directDescendants := lf.DirectDescendantsAfter(day)
	key := fmt.Sprintf("(%d,%d)", lf.SpawnTimer, day)

	if n, ok := m[key]; ok {
		return n
	}

	sum := 0
	spawnDay := day

	for i := 0; i < directDescendants; i++ {
		if i == 0 {
			spawnDay -= (lf.SpawnTimer + 1)
		} else {
			spawnDay -= 7
		}
		child := NewLanternFish(8)
		sum = sum + 1 + child.AllDescendantsAfter(spawnDay, m)
	}

	m[key] = sum

	return sum
}

func (lf *LanternFish) DirectDescendantsAfter(day int) int {
	actualDays := day - (lf.SpawnTimer + 1)

	// we cannot spawn within the amount of time
	if actualDays < 0 {
		return 0
	}

	// we can spawn at least once. how many direct descendants?
	remaining := (actualDays / 7)
	sum := 1 + remaining

	return sum
}

func (lf *LanternFish) String() string {
	return fmt.Sprintf("%d", lf.SpawnTimer)
}

func SchoolSizeAfter(day int, school []*LanternFish) int {
	size := 0
	m := make(map[string]int)

	for _, lf := range school {
		size = size + 1 + lf.AllDescendantsAfter(day, m)
	}

	return size
}

func part1() {
	school := load()

	maxDays := 80
	day := 1

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

		day++
	}

	fmt.Println("part 1: ", len(school))
}

func part2() {
	school := load()

	fmt.Println("part 2: ", SchoolSizeAfter(256, school))
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
