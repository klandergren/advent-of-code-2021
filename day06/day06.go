package day06

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartOne(reader io.Reader) (int, error) {
	school, err := load(reader)

	if err != nil {
		return -1, err
	}

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

	return len(school), nil
}
func PartTwo(reader io.Reader) (int, error) {
	school, err := load(reader)

	if err != nil {
		return -1, err
	}

	return SchoolSizeAfter(256, school), nil
}

func SchoolSizeAfter(day int, school []*LanternFish) int {
	size := 0
	m := make(map[string]int)

	for _, lf := range school {
		size = size + 1 + lf.AllDescendantsAfter(day, m)
	}

	return size
}

func load(reader io.Reader) (school []*LanternFish, err error) {
	bReader := bufio.NewReader(reader)

	school = make([]*LanternFish, 0)
	for {
		line, _, err := bReader.ReadLine()

		if err == io.EOF {
			break
		}

		spawnTimers := strings.Split(string(line), ",")

		for _, spawnTimer := range spawnTimers {
			st, err := strconv.Atoi(spawnTimer)

			if err != nil {
				return nil, err
			}

			school = append(school, NewLanternFish(st))
		}

	}

	return school, nil
}
