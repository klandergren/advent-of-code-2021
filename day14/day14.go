package day14

import (
	"io"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	p, err := Load(reader)
	if err != nil {
		return -1, err
	}

	p.ApplySteps(10)

	return p.Answer(), nil
}

func PartTwo(reader io.Reader) (int, error) {
	p, err := Load(reader)
	if err != nil {
		return -1, err
	}

	p.ApplySteps(40)

	return p.Answer(), nil
}

func Load(reader io.Reader) (p *Polymer, err error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return nil, err
	}

	var templateLines []string
	var ruleLines []string
	seenBlank := false
	for _, line := range lines {
		if line == "" {
			seenBlank = true
			continue
		}

		if seenBlank {
			ruleLines = append(ruleLines, line)
		} else {
			templateLines = append(templateLines, line)
		}
	}

	template := []rune(templateLines[0])

	rules := make([]Rule, 0)
	for _, line := range ruleLines {
		rule := NewRule(line)
		rules = append(rules, rule)
	}

	return NewPolymer(template, rules), nil
}
