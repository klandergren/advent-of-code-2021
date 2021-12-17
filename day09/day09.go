package day09

import (
	"io"
	"sort"

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
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	hm, err := NewHeightMap(lines)
	if err != nil {
		return -1, err
	}

	bs := hm.BasinSizes()

	sort.Sort(sort.Reverse(sort.IntSlice(bs)))

	product := 1
	for i := 0; i < 3; i++ {
		product *= bs[i]
	}

	return product, nil
}
