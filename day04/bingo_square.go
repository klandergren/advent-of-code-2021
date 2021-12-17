package day04

import "fmt"

type BingoSquare struct {
	Value    int
	IsMarked bool
}

func (sq *BingoSquare) String() string {
	if sq.IsMarked {
		return fmt.Sprintf("%2d'", sq.Value)
	} else {
		return fmt.Sprintf("%2d ", sq.Value)
	}
}
