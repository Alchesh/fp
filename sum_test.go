package main

import "testing"

var m = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func BenchmarkSum50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(m)
	}
}

func BenchmarkSumR50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumR(m)
	}
}
