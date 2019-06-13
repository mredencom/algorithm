package recursion

import "fmt"

//递归
//阶乘的使用
//传统的方法迭代实现阶乘
type Fac struct {
	val map[int]int
}

//建立一个阶乘
func NewFactorial(n int) *Fac {
	return &Fac{make(map[int]int, n)}
}

//实现阶乘
func (f *Fac) Factorial(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}
	//查找n>1情况
	if n > 1 {
		result := n * f.Factorial(n-1)
		f.val[n] = result
		return result
	} else {
		//n<=1情况
		f.val[n] = 1
		return 1
	}

}

//打印出阶乘
func (f *Fac) Print(n int) {
	fmt.Println(f.val[n])
}
