package day07

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func PartOne(reader io.Reader) (int, error) {
	horizontalPositions, err := load(reader)
	if err != nil {
		return -1, err
	}

	minFuelCost, err := FindMinFuelCostPositionPartOneBrute(horizontalPositions)
	if err != nil {
		return -1, err
	}

	return minFuelCost, nil
}
func PartTwo(reader io.Reader) (int, error) {
	horizontalPositions, err := load(reader)
	if err != nil {
		return -1, err
	}

	minFuelCost, err := FindMinFuelCostPositionPartTwoBrute(horizontalPositions)
	if err != nil {
		return -1, err
	}

	return minFuelCost, nil
}

func MinMax(s []int) (min int, max int, err error) {
	if s == nil {
		return -1, -1, errors.New("nil slice")
	}

	if len(s) == 0 {
		return -1, -1, errors.New("empty slice")
	}

	min = s[0]
	max = s[0]

	for i := 1; i < len(s); i++ {
		maybeMin := s[i]
		maybeMax := s[i]

		if maybeMin < min {
			min = maybeMin
		}

		if max < maybeMax {
			max = maybeMax
		}

	}

	return min, max, nil
}

func FuelCostPartOne(positions []int, alignmentPosition int) (int, error) {
	if positions == nil {
		return -1, errors.New("nil slice")
	}

	if len(positions) == 0 {
		return -1, errors.New("empty slice")
	}

	min, max, err := MinMax(positions)
	if err != nil {
		return -1, err
	}

	if alignmentPosition < min {
		return -1, errors.New("alignment position too low")
	}

	if max < alignmentPosition {
		return -1, errors.New("alignment position too low")
	}

	cost := 0

	for _, p := range positions {
		if p <= alignmentPosition {
			cost = cost + (alignmentPosition - p)
		} else {
			cost = cost + (p - alignmentPosition)
		}
	}

	return cost, nil
}

func FuelCostPartTwo(positions []int, alignmentPosition int) (int, error) {
	if positions == nil {
		return -1, errors.New("nil slice")
	}

	if len(positions) == 0 {
		return -1, errors.New("empty slice")
	}

	min, max, err := MinMax(positions)
	if err != nil {
		return -1, err
	}

	if alignmentPosition < min {
		return -1, errors.New("alignment position too low")
	}

	if max < alignmentPosition {
		return -1, errors.New("alignment position too low")
	}

	cost := 0

	for _, p := range positions {
		if p <= alignmentPosition {
			cost = cost + FuelCostToMove(alignmentPosition-p)
		} else {
			cost = cost + FuelCostToMove(p-alignmentPosition)
		}
	}

	return cost, nil
}

func FuelCostToMove(numPositions int) int {
	cost := 0
	for i := 1; i <= numPositions; i++ {
		cost = cost + i
	}

	return cost
}

func FindMinFuelCostPositionPartTwoBrute(positions []int) (int, error) {
	min, max, err := MinMax(positions)
	if err != nil {
		return -1, err
	}

	minFuelCost, err := FuelCostPartTwo(positions, min)
	if err != nil {
		return -1, err
	}

	for i := min + 1; i <= max; i++ {
		maybeMinFuelCost, err := FuelCostPartTwo(positions, i)
		if err != nil {
			return -1, err
		}
		if maybeMinFuelCost < minFuelCost {
			minFuelCost = maybeMinFuelCost
		}
	}

	return minFuelCost, nil
}

func FindMinFuelCostPositionPartOneBrute(positions []int) (int, error) {
	min, max, err := MinMax(positions)
	if err != nil {
		return -1, err
	}

	minFuelCost, err := FuelCostPartOne(positions, min)
	if err != nil {
		return -1, err
	}

	for i := min + 1; i <= max; i++ {
		maybeMinFuelCost, err := FuelCostPartOne(positions, i)
		if err != nil {
			return -1, err
		}
		if maybeMinFuelCost < minFuelCost {
			minFuelCost = maybeMinFuelCost
		}
	}

	return minFuelCost, nil
}

func load(reader io.Reader) (horizontalPositions []int, err error) {
	bReader := bufio.NewReader(reader)

	for {
		line, _, err := bReader.ReadLine()

		if err == io.EOF {
			break
		}

		for _, sPos := range strings.Split(string(line), ",") {
			pos, err := strconv.Atoi(sPos)

			if err != nil {
				return nil, err
			}

			horizontalPositions = append(horizontalPositions, pos)
		}

	}

	return horizontalPositions, nil
}
