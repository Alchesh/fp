package main

import "fmt"

func main() {

	val := 4

	fmt.Println("Fib = ", Fib(val), "\nCahedFib = ", FibCashed(val))
	fmt.Println("\nAll Done!")
}
