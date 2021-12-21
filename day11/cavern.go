package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Cavern struct {
	GridYX          [][]*Octopus
	TotalFlashCount int
	StepCount       int
}

func NewCavern(lines []string) (c *Cavern, err error) {
	size := len(lines)

	gridYX := make([][]*Octopus, size)
	for y, _ := range gridYX {
		gridYX[y] = make([]*Octopus, size)
	}

	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, c := range chars {
			energy, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			gridYX[y][x] = &Octopus{
				EnergyLevel: energy,
			}
		}
	}

	return &Cavern{GridYX: gridYX}, nil
}

func (c *Cavern) FirstSynchronizeStep() (step int) {
	numOctopi := len(c.GridYX) * len(c.GridYX)
	numZeros := 0
	for ok := true; ok; ok = (numZeros != numOctopi) {
		numZeros = 0
		c.Step()
		for _, row := range c.GridYX {
			for _, o := range row {
				if o.EnergyLevel == 0 {
					numZeros += 1
				}
			}
		}
	}
	return c.StepCount
}

// just the amount of flashes during `steps`; inspect `TotalFlashCount` for total
func (c *Cavern) Advance(steps int) (flashCount int) {
	for i := 0; i < steps; i++ {
		flashCount += c.Step()
	}
	return flashCount
}

func (c *Cavern) Step() (flashCount int) {
	// increment all by 1
	for _, row := range c.GridYX {
		for _, o := range row {
			o.Increment()
		}
	}

	// loop through. if can flash, flash it and increment energy of nearby.
	// continue looping until CanFlash() returns all false
	var flashes int
	for ok := true; ok; ok = (0 < flashes) {
		flashes = 0
		for y, row := range c.GridYX {
			for x, o := range row {
				if o.CanFlash() {
					o.Flash()
					c.Flash(x, y)
					flashes++
				}
			}
		}
		flashCount += flashes
	}

	// call `SmartReset on all
	for _, row := range c.GridYX {
		for _, o := range row {
			o.SmartReset()
		}
	}

	// increment step count
	c.StepCount += 1

	// increment total flash count
	c.TotalFlashCount += flashCount

	return flashCount
}

func (c *Cavern) Flash(x, y int) {
	for _, o := range c.Adjacent(x, y) {
		o.Increment()
	}
}

func (c *Cavern) Adjacent(x, y int) (adj []*Octopus) {
	// north
	if o := c.Get(x, y-1); o != nil {
		adj = append(adj, o)
	}

	// north east
	if o := c.Get(x+1, y-1); o != nil {
		adj = append(adj, o)
	}

	// east
	if o := c.Get(x+1, y); o != nil {
		adj = append(adj, o)
	}

	// south east
	if o := c.Get(x+1, y+1); o != nil {
		adj = append(adj, o)
	}

	// south
	if o := c.Get(x, y+1); o != nil {
		adj = append(adj, o)
	}

	// south west
	if o := c.Get(x-1, y+1); o != nil {
		adj = append(adj, o)
	}

	// west
	if o := c.Get(x-1, y); o != nil {
		adj = append(adj, o)
	}

	// north west
	if o := c.Get(x-1, y-1); o != nil {
		adj = append(adj, o)
	}

	return adj
}

func (c *Cavern) Get(x, y int) (o *Octopus) {
	if y < 0 {
		return nil
	}

	if len(c.GridYX) <= y {
		return nil
	}

	if x < 0 {
		return nil
	}

	if len(c.GridYX[y]) <= x {
		return nil
	}

	return c.GridYX[y][x]
}

func (c *Cavern) String() string {
	var sb strings.Builder
	for _, row := range c.GridYX {
		for _, o := range row {
			sb.WriteString(fmt.Sprintf("%s", o))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
