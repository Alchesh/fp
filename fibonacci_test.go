package main

import "testing"

var Tests = []struct {
	val      int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{20, 6765},
	{42, 267914296},
}

func TestFib(t *testing.T) {
	for _, test := range Tests {
		if v := Fib(test.val); v != test.expected {
			t.Errorf("Fib(%d) returned %d, expected %d", test.val, v, test.expected)
		}
	}
}

func TestFibCached(t *testing.T) {
	for _, test := range Tests {
		if v := FibCashed(test.val); v != test.expected {
			t.Errorf("FibCached(%d) returned %d, expected %d", test.val, v, test.expected)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

func BenchmarkFibCahed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibCashed(20)
	}
}
