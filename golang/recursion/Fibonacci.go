package recursion

import "fmt"

//递归的一种
//实现斐波那契数列
type Fibs struct {
	val map[int]int
}

//建立一个斐波那契
func NewFibonacci(n int) *Fibs {
	return &Fibs{make(map[int]int, n)}
}

//实现菲波那契数列
//公式: F(1)=1，F(2)=1, F(n)=F(n-1)+F(n-2)（n>=3，n∈N*） 看这个这个递推公式很容易就可以实现程序
func (fib *Fibs) Fibonacci(n int) int {
	if fib.val[n] != 0 {
		return fib.val[n]
	}
	if n <= 1 {
		fib.val[n] = 1
		return 1
	} else if n == 2 {
		fib.val[2] = 1
		return 1
	} else {
		result := fib.val[n-1] + fib.val[n-2]
		fib.val[n] = result
		return result
	}
}

//打印数列
func (fib *Fibs) Print(n int) {
	fmt.Println(fib.val[n])
}
