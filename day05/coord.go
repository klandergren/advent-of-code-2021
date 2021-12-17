package day05

import "fmt"

type Coord struct {
	X, Y int
}

func (c *Coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
