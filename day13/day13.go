package day13

import (
	"fmt"
	"io"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	tp, instructions, err := Load(reader)
	if err != nil {
		return -1, err
	}

	tp.Apply(instructions[0])

	return tp.DotsVisible(), nil
}

func PartTwo(reader io.Reader) (int, error) {
	tp, instructions, err := Load(reader)
	if err != nil {
		return -1, err
	}

	for _, i := range instructions {
		tp.Apply(i)
	}

	fmt.Println(tp)

	return -1, nil
}

func Load(reader io.Reader) (tp *TransparentPaper, instructions []Instruction, err error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return nil, instructions, err
	}

	var tpLines []string
	var iLines []string
	seenBlank := false
	for _, line := range lines {
		if line == "" {
			seenBlank = true
			continue
		}

		if seenBlank {
			iLines = append(iLines, line)
		} else {
			tpLines = append(tpLines, line)
		}
	}

	tp, err = LoadTransparentPaper(tpLines)
	if err != nil {
		return nil, instructions, err
	}

	instructions, err = LoadInstructions(iLines)
	if err != nil {
		return nil, instructions, err
	}

	return tp, instructions, nil
}
