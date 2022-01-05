package day14

import "testing"

func Test_RuleGeneration(t *testing.T) {
	testCases := map[string]struct {
		line    string
		match   []rune
		element rune
	}{
		"CH -> B": {
			line:    "CH -> B",
			match:   []rune{'C', 'H'},
			element: 'B',
		},
		"HH -> N": {
			line:    "HH -> N",
			match:   []rune{'H', 'H'},
			element: 'N',
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(tSub *testing.T) {

			tSub.Parallel()

			r := NewRule(tc.line)

			if !r.Matches(tc.match) {
				tSub.Fatalf("expected true but got: %t\n", r.Matches(tc.match))
			}

			if r.Element() != tc.element {
				tSub.Fatalf("expected %q but got: %q\n", tc.element, r.Element())
			}

		})
	}
}
