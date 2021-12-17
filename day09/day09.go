package day09

import (
	"errors"
	"io"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	hm, err := NewHeightMap(lines)
	if err != nil {
		return -1, err
	}

	return hm.SumRisk(), nil
}

func PartTwo(reader io.Reader) (int, error) {
	return -1, errors.New("not implemented yet")
}
