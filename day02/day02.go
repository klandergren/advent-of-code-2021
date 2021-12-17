package day02

import (
	"io"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	sub := Submarine{0, 0, 0}
	for _, line := range lines {
		sub.AdvanceWithoutAim(line)
	}

	return sub.Mult(), nil
}

func PartTwo(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	sub := Submarine{0, 0, 0}
	for _, line := range lines {
		sub.AdvanceWithAim(line)
	}

	return sub.Mult(), nil
}
