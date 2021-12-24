package day12

import "unicode"

type Cave struct {
	Name      string
	IsBigCave bool
}

func NewCave(name string) Cave {
	allLowerCase := true
	for _, r := range name {
		allLowerCase = allLowerCase && unicode.IsLower(r)
	}

	if allLowerCase {
		return Cave{Name: name, IsBigCave: false}
	} else {
		return Cave{Name: name, IsBigCave: true}
	}
}

func (c Cave) String() string {
	return c.Name
}
