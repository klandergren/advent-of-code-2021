package day13

import "testing"

func Test_New(t *testing.T) {
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
