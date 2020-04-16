package main

import "fmt"

func main() {

	val := 7

	fmt.Println("Fib = ", Fib(val), "\nCahedFib = ", FibCashed(val))
	fmt.Println("\nAll Done!")
}
