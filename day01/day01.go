package day01

import (
	"io"
	"math"
	"strconv"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	prev_depth := math.MaxInt64
	c := 0
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}

		if prev_depth < depth {
			c++
		}

		prev_depth = depth
	}

	return c, nil
}

func PartTwo(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	prev_window := math.MaxInt64
	depths := make([]int, 0)
	c := 0
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}

		depths = append(depths, depth)

		if len(depths) < 3 {
			// not enough measurements yet
			continue
		} else if len(depths) == 3 {
			// we have first measurement
		} else {
			// pop
			_, depths = depths[0], depths[1:]
		}

		window := 0
		for _, d := range depths {
			window += d
		}

		if prev_window < window {
			c++
		}

		prev_window = window
	}

	return c, nil
}
