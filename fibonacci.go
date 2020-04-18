package main

/*
import (
	"fmt"
)
*/

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

// FibCashed - fibonacci with cache
func FibCashed(val int) int {

	FibCash = Cache(func(val int) int {
		if val == 0 {
			return 0
		} else if val <= 2 {
			return 1
		}

		return FibCash(val-2) + FibCash(val-1)
	})
	return FibCash(val)
}

func Channel(ch chan int, counter int) {
	n1, n2 := 0, 1
	for i := 0; i < counter; i++ {
		ch <- n1
		n1, n2 = n2, n1+n2
	}
	close(ch)
}

// FibCashed - fibonacci with channels
func FibChanneled(n int) int {
	var result int

	n += 2
	ch := make(chan int)
	go Channel(ch, n-1) // one channel for one number
	i := 0

	for num := range ch {
		result = num
		//fmt.Println("result =", result)
		i++
	}
	return result
}
