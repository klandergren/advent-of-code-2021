package day05

import "fmt"

type VentLine struct {
	Start, End *Coord
}

func (vl *VentLine) IsDiagonal() bool {
	return vl.Start.X != vl.End.X && vl.Start.Y != vl.End.Y
}

func (vl *VentLine) MaxX() int {
	if vl.Start.X < vl.End.X {
		return vl.End.X
	} else {
		return vl.Start.X
	}
}

func (vl *VentLine) MaxY() int {
	if vl.Start.Y < vl.End.Y {
		return vl.End.Y
	} else {
		return vl.Start.Y
	}
}

// assume lines only cross whole int coordinates
func (vl *VentLine) Coords() []*Coord {
	coords := make([]*Coord, 0)

	x, y := vl.Start.X, vl.Start.Y

	if vl.IsDiagonal() {
		switch {
		case vl.Start.X <= vl.End.X && vl.Start.Y <= vl.End.Y:
			for x <= vl.End.X && y <= vl.End.Y {
				coords = append(coords, &Coord{x, y})
				x++
				y++
			}
		case vl.Start.X <= vl.End.X && vl.End.Y <= vl.Start.Y:
			for x <= vl.End.X && vl.End.Y <= y {
				coords = append(coords, &Coord{x, y})
				x++
				y--
			}
		case vl.End.X <= vl.Start.X && vl.Start.Y <= vl.End.Y:
			for vl.End.X <= x && y <= vl.End.Y {
				coords = append(coords, &Coord{x, y})
				x--
				y++
			}
		case vl.End.X <= vl.Start.X && vl.End.Y <= vl.Start.Y:
			for vl.End.X <= x && vl.End.Y <= y {
				coords = append(coords, &Coord{x, y})
				x--
				y--
			}
		}
	} else {
		switch {
		case vl.Start.X == vl.End.X:
			if vl.Start.Y < vl.End.Y {
				for y <= vl.End.Y {
					coords = append(coords, &Coord{x, y})
					y++
				}
			} else {
				for vl.End.Y <= y {
					coords = append(coords, &Coord{x, y})
					y--
				}
			}
		case vl.Start.Y == vl.End.Y:
			if vl.Start.X < vl.End.X {
				for x <= vl.End.X {
					coords = append(coords, &Coord{x, y})
					x++
				}
			} else {
				for vl.End.X <= x {
					coords = append(coords, &Coord{x, y})
					x--
				}
			}
		}
	}

	return coords
}

func (vl *VentLine) String() string {
	return fmt.Sprintf("%s -> %s", vl.Start, vl.End)
}
