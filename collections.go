package main

//import "fmt"

const INVALID_INT = -1
const INVALID_STRING = ""

type Iterator interface {
	Next() (val string, ok bool)
}

type Collection struct {
	i   int
	mas []string
}

func (c *Collection) Next() (val string, ok bool) {
	c.i++
	if c.i >= len(c.mas) {
		return INVALID_STRING, false
	}

	return c.mas[c.i], true
}

func newSlice(s []string) *Collection {
	return &Collection{INVALID_INT, s}
}
