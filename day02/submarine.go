package day02

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Submarine struct {
	Aim, X, Depth int
}

func (sub *Submarine) Mult() int {
	return sub.X * sub.Depth
}

func (sub *Submarine) AdvanceWithoutAim(step string) error {
	parts := strings.Fields(step)

	if len(parts) != 2 {
		return fmt.Errorf("unexpected step format: %s", step)
	}

	direction := parts[0]
	magnitude, err := strconv.Atoi(parts[1])

	if err != nil {
		return err
	}

	switch direction {
	case "forward":
		sub.X = sub.X + magnitude
	case "down":
		sub.Depth = sub.Depth + magnitude
	case "up":
		sub.Depth = sub.Depth - magnitude
	default:
		return errors.New("how'd we get here?")
	}

	return nil
}

func (sub *Submarine) AdvanceWithAim(step string) error {
	parts := strings.Fields(step)

	if len(parts) != 2 {
		return fmt.Errorf("unexpected step format: %s", step)
	}

	direction := parts[0]
	magnitude, err := strconv.Atoi(parts[1])

	if err != nil {
		return err
	}

	switch direction {
	case "forward":
		sub.X = sub.X + magnitude
		sub.Depth = sub.Depth + sub.Aim*magnitude
	case "down":
		sub.Aim = sub.Aim + magnitude
	case "up":
		sub.Aim = sub.Aim - magnitude
	default:
		return errors.New("how'd we get here?")
	}

	return nil
}
