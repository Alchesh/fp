package main

/*
import (
	"fmt"
)
*/

func Sum(val []int) int {

	l := len(val)
	s := 0

	for i := 0; i < l; i++ {
		s += val[i]
	}

	return s
}

func SumR(val []int) int {

	if len(val) == 0 {
		return 0
	}

	return val[0] + SumR(val[1:])
}
