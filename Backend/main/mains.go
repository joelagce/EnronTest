package main

import "testing"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return (Fibonacci(n-1) + Fibonacci(n-2))
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		n    int
		fibo int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{15, 610},
		{17, 1597},
		{40, 102334155},
	}

	for _, table := range tables {
		result := Fibonacci(table.n)
		if result != table.fibo {
			t.Errorf("Fibonacci incorrecta, esperabamos %d, pero obtubimos %d", table.fibo, result)
		}
	}
}
