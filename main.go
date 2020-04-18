package main

import "fmt"

func main() {

	/*	val := 7
		fmt.Println("Fib = ", Fib(val))
		FibCash = FibCashed
		fmt.Println("\nCahedFib = ", FibCashed(val))
		fmt.Println("\nChanneledFib = ", FibChanneled1(val))
	*/
	/*
		m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println("Sum = ", Sum(m))
		fmt.Println("Sum = ", SumR(m))
	*/

	var intCollection Iterator
	intCollection = newSlice([]string{"CRV", "IS250", "Blazer"})
	value, ok := intCollection.Next()
	for ok {
		//println(value)
		fmt.Println(value)
		value, ok = intCollection.Next()
	}

	fmt.Println("\nAll Done!")
}
