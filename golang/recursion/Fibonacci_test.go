package recursion

import "testing"

//测试斐波那契
//go test -v -run TestFibs_Fibonacci -o Fibonacci_test.go
func TestFibs_Fibonacci(t *testing.T) {
	fib := NewFibonacci(10)
	for i := 1; i < 15; i++ {
		fib.Fibonacci(i)
		fib.Print(i)
	}
}
