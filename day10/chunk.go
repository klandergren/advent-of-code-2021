package day10

import "fmt"

type Chunk struct {
	OpenChar string
	Complete bool
	Chunks   []Chunk
}

func (c *Chunk) String() string {
	return fmt.Sprintf("'%s'", c.OpenChar)
}
