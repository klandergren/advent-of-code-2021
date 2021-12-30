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

func NewInstruction(line string) (*Instruction, error) {
	parts := strings.Split(line, "=")

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	var direction string
	if strings.Contains(line, "x") {
		direction = "left"
	} else {
		direction = "up"
	}

	return &Instruction{
		Raw:       line,
		Direction: direction,
		Value:     value,
	}, nil
}
