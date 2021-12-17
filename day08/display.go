package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type Display struct {
	InputSignals []string
}

func (d *Display) OutputValue() (int, error) {
	// create digits
	digits := make([]*Digit, len(d.InputSignals))

	for i, s := range d.InputSignals {
		d := NewDigitFromInput(s)
		digits[i] = d
	}

	// create string of values
	var sb strings.Builder
	for _, d := range digits {
		sb.WriteString(fmt.Sprintf("%d", d.Value()))
	}

	v, err := strconv.Atoi(sb.String())

	if err != nil {
		return -1, err
	}

	return v, nil
}
