package day12

import (
	"io"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		return -1, err
	}

	shouldAllowDouble := false
	numAllPaths := len(cn.AllPaths(shouldAllowDouble))

	return numAllPaths, nil
}

func PartTwo(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	cn, err := NewCaveNetwork(lines)
	if err != nil {
		return -1, err
	}

	shouldAllowDouble := true
	numAllPaths := len(cn.AllPaths(shouldAllowDouble))

	return numAllPaths, nil
}
