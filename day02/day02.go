package day02

import (
	"bufio"
	"io"
)

func PartOne(reader *bufio.Reader) (int, error) {
	sub := Submarine{0, 0, 0}

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		sub.AdvanceWithoutAim(string(line))
	}

	return sub.Mult(), nil
}

func PartTwo(reader *bufio.Reader) (int, error) {
	sub := Submarine{0, 0, 0}

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		sub.AdvanceWithAim(string(line))
	}

	return sub.Mult(), nil
}
