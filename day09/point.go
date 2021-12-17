package day09

import "fmt"

type Point struct {
	X, Y, Height int
}

func (p *Point) Risk() int {
	return 1 + p.Height
}

func (p *Point) String() string {
	return fmt.Sprintf("x: %d, y: %d - height: %d", p.X, p.Y, p.Height)
}
