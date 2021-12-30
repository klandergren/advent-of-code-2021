package day13

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Raw       string
	Direction string
	Value     int
}

func LoadInstructions(lines []string) (instructions []Instruction, err error) {
	for _, line := range lines {
		i, err := NewInstruction(line)
		if err != nil {
			return instructions, err
		}

		instructions = append(instructions, i)
	}
	return instructions, nil
}

func NewInstruction(line string) (Instruction, error) {
	parts := strings.Split(line, "=")

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return Instruction{}, err
	}

	var direction string
	if strings.Contains(line, "x") {
		direction = "left"
	} else {
		direction = "up"
	}

	return Instruction{
		Raw:       line,
		Direction: direction,
		Value:     value,
	}, nil
}
