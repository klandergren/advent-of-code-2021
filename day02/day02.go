package day02

import (
	"bufio"
	"io"
)

func PartOne(reader io.Reader) (int, error) {
	sub := Submarine{0, 0, 0}

	bReader := bufio.NewReader(reader)

	for {
		line, _, err := bReader.ReadLine()

		if err == io.EOF {
			break
		}

		sub.AdvanceWithoutAim(string(line))
	}

	return sub.Mult(), nil
}

func PartTwo(reader io.Reader) (int, error) {
	sub := Submarine{0, 0, 0}

	bReader := bufio.NewReader(reader)

	for {
		line, _, err := bReader.ReadLine()

		if err == io.EOF {
			break
		}

		sub.AdvanceWithAim(string(line))
	}

	return sub.Mult(), nil
}
