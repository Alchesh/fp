package main

//import "fmt"

type CachedFunc func(int) int

var FibCash = Cache(Fib) // Init cache on start app

// Fib - fibonacci
// 1 1 2 3 5 8 13
func Fib(val int) int {

	if val == 0 {
		return 0
	} else if val <= 2 {
		return 1
	}

	return Fib(val-1) + Fib(val-2)
}

func Cache(m CachedFunc) CachedFunc {

	c := make(map[int]int)

	return func(key int) int {
		if val, found := c[key]; found {
			return val
		}

		temp := m(key)
		c[key] = temp

		return temp
	}
}

func FibCashed(val int) int {
	return FibCash(val)
}
