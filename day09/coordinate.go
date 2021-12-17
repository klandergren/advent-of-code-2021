package day09

import "fmt"

type Coordinate struct {
	X, Y int
}

func (c *Coordinate) AdjacentCoordinates() (adj []*Coordinate) {
	adj = append(adj, c.North())
	adj = append(adj, c.West())
	adj = append(adj, c.South())
	adj = append(adj, c.East())

	return adj
}

func (c *Coordinate) North() *Coordinate {
	return &Coordinate{
		X: c.X,
		Y: c.Y - 1,
	}
}

func (c *Coordinate) West() *Coordinate {
	return &Coordinate{
		X: c.X - 1,
		Y: c.Y,
	}
}

func (c *Coordinate) South() *Coordinate {
	return &Coordinate{
		X: c.X,
		Y: c.Y + 1,
	}
}

func (c *Coordinate) East() *Coordinate {
	return &Coordinate{
		X: c.X + 1,
		Y: c.Y,
	}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}
