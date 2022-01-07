package day14

import "sort"

type Polymer struct {
	Template  []rune
	Rules     []Rule
	Sequence  []rune
	Step      int
	Histogram map[rune]int
}

func NewPolymer(template []rune, rules []Rule) *Polymer {
	h := make(map[rune]int)

	for _, r := range template {
		h[r] += 1
	}

	return &Polymer{
		Template:  template,
		Rules:     rules,
		Sequence:  template,
		Step:      0,
		Histogram: h,
	}
}

func (p *Polymer) ApplySteps2(n int) {
	m := make(map[string](map[rune]int), 0)

}

func (p *Polymer) ApplySteps(n int) {
	for i := 0; i < n; i++ {
		newSequence := make([]rune, 0)

		for i, j := 0, 1; j < len(p.Sequence); i, j = i+1, j+1 {
			r1 := p.Sequence[i]
			r2 := p.Sequence[j]

			pair := p.Sequence[i:(j + 1)]

			// always append r1 before applying rule matching
			newSequence = append(newSequence, r1)
			for _, rule := range p.Rules {
				if rule.Matches(pair) {
					e := rule.Element()
					newSequence = append(newSequence, e)
					p.Histogram[e] += 1
					// we know there will only be one matching rule
					break
				}
			}

			// only apply r2 if at the end
			if j+1 == len(p.Sequence) {
				newSequence = append(newSequence, r2)
			}
		}

		p.Sequence = newSequence
		p.Step += 1
	}
}

func (p *Polymer) Answer() int {
	return p.MostCommon() - p.LeastCommon()
}

func (p *Polymer) MostCommon() int {
	values := make([]int, 0)
	for _, v := range p.Histogram {
		values = append(values, v)
	}

	sort.Ints(values)

	return values[len(values)-1]
}

func (p *Polymer) LeastCommon() int {
	values := make([]int, 0)
	for _, v := range p.Histogram {
		values = append(values, v)
	}

	sort.Ints(values)

	return values[0]
}

func (p *Polymer) CountOf(r rune) int {
	return p.Histogram[r]
}

func (p *Polymer) Length() int {
	return len(p.Sequence)
}
