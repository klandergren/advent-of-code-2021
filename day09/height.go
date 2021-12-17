package day09

type Height int

func (h Height) Risk() int {
	return int(h) + 1
}
