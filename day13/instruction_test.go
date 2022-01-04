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
	testCases := map[string]struct {
		line        string
		instruction Instruction
	}{
		"y direction": {
			line: "fold along y=7",
			instruction: Instruction{
				Raw:       "fold along y=7",
				Direction: "up",
				Value:     7,
			},
		},
		"x direction": {
			line: "fold along x=5",
			instruction: Instruction{
				Raw:       "fold along x=5",
				Direction: "left",
				Value:     5,
			},
		},
	}

	for name, tc := range testCases {
		testCase := tc // capture for scope
		t.Run(name, func(tSub *testing.T) {
			tSub.Parallel()

			i, err := NewInstruction(testCase.line)
			if err != nil {
				tSub.Fatal(err)
			}

			if i.Raw != testCase.instruction.Raw {
				tSub.Fatalf("instruction should have Raw: %q got: %q\n", testCase.instruction.Raw, i.Raw)
			}

			if i.Direction != testCase.instruction.Direction {
				tSub.Fatalf("instruction should have Direction: %q got: %q\n", testCase.instruction.Direction, i.Direction)
			}

			if i.Value != testCase.instruction.Value {
				tSub.Fatalf("instruction should have Value: %d got: %d\n", testCase.instruction.Value, i.Value)
			}

		})

	}

}
