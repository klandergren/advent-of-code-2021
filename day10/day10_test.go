package day10

import (
	"strings"
	"testing"
)

var testInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(testInput)

	score, err := PartOne(r)
	if err != nil {
		t.Error(err)
	}

	if score != 26397 {
		t.Errorf("score for test input should be 26397, got: %d", score)
	}

}

func Test_PartTwo(t *testing.T) {
	t.SkipNow()
}
