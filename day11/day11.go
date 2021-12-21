package day11

import (
	"io"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	c, err := NewCavern(lines)
	if err != nil {
		return -1, err
	}

	flashCount := c.Advance(100)

	return flashCount, nil
}

func PartTwo(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	c, err := NewCavern(lines)
	if err != nil {
		return -1, err
	}

	firstSynchronizeStep := c.FirstSynchronizeStep()

	return firstSynchronizeStep, nil
}
