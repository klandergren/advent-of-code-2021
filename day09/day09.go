package day09

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func PartOne(reader *bufio.Reader) (int, error) {
	return -1, errors.New("not implemented yet")
}
func PartTwo(reader *bufio.Reader) (int, error) {
	return -1, errors.New("not implemented yet")
}

func SumRisk(points []*Point) (sumRisk int) {
	for _, p := range points {
		sumRisk += p.Risk()
	}
	return sumRisk
}

func AdjacentPoints(point *Point, hm [][]*Point) (adjacentPoints []*Point) {
	return nil
}

func loadLines(reader io.Reader) (lines []string, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func load(reader io.Reader) (hm [][]*Point, err error) {
	lines, err := loadLines(reader)
	if err != nil {
		return nil, err
	}

	for y, line := range lines {
		rowData := strings.Split(line, "")
		hm = append(hm, make([]*Point, len(rowData)))
		for x, sHeight := range rowData {
			nHeight, err := strconv.Atoi(sHeight)
			if err != nil {
				return nil, err
			}

			hm[y][x] = &Point{x, y, nHeight}
		}
	}

	return hm, nil
}
