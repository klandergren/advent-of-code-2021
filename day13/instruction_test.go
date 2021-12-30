package day13

import (
	"strings"
	"testing"

	"github.com/klandergren/advent-of-code-2021/util"
)

func Test_LoadInstructions(t *testing.T) {
	data := `fold along x=655
fold along y=447
fold along x=327
fold along y=223
fold along x=163
fold along y=111
fold along x=81
fold along y=55
fold along x=40
fold along y=27
fold along y=13
fold along y=6
`

	r := strings.NewReader(data)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	instructions, err := LoadInstructions(lines)
	if err != nil {
		t.Error(err)
	}

	if len(instructions) != 12 {
		t.Errorf("instructions should have length 12, got: %d\n", len(instructions))
	}

}

func Test_NewInstructions(t *testing.T) {
	data := []struct {
		line        string
		instruction Instruction
	}{
		{
			line: "fold along y=7",
			instruction: Instruction{
				Raw:       "fold along y=7",
				Direction: "up",
				Value:     7,
			},
		},
		{
			line: "fold along x=5",
			instruction: Instruction{
				Raw:       "fold along x=5",
				Direction: "left",
				Value:     5,
			},
		},
	}

	for _, expected := range data {
		i, err := NewInstruction(expected.line)
		if err != nil {
			t.Error(err)
		}

		if i.Raw != expected.instruction.Raw {
			t.Errorf("instruction should have Raw: %q got: %q\n", expected.instruction.Raw, i.Raw)
		}

		if i.Direction != expected.instruction.Direction {
			t.Errorf("instruction should have Direction: %q got: %q\n", expected.instruction.Direction, i.Direction)
		}

		if i.Value != expected.instruction.Value {
			t.Errorf("instruction should have Value: %d got: %d\n", expected.instruction.Value, i.Value)
		}

	}

}
