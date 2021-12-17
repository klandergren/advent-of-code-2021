package day01

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

func PartOne(reader io.Reader) (int, error) {
	prev_depth := math.MaxInt64

	c := 0

	bReader := bufio.NewReader(reader)

	for {
		l, _, err := bReader.ReadLine()

		if err == io.EOF {
			break
		}

		depth, _ := strconv.Atoi(string(l))

		if prev_depth < depth {
			c++
		}

		prev_depth = depth
	}

	return c, nil
}

func PartTwo(reader io.Reader) (int, error) {
	prev_window := math.MaxInt64
	depths := make([]int, 0)

	c := 0

	bReader := bufio.NewReader(reader)

	for {
		l, _, err := bReader.ReadLine()

		if err == io.EOF {
			break
		}

		depth, _ := strconv.Atoi(string(l))

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
