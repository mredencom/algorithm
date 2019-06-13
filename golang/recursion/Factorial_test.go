package recursion

import "testing"

//测试EnQueue 方法
//go test -v -run TestFac_Factorial -o Factorial_test.go
/** 结果
=== RUN   TestFac_Factorial
		1
		2
		6
		24
		120
		720
		5040
		40320
		362880
		3628800
		39916800
		479001600
		6227020800
		87178291200
		1307674368000
		20922789888000
		355687428096000
		6402373705728000
		121645100408832000
*/
func TestFac_Factorial(t *testing.T) {
	f := NewFactorial(10)
	for i := 1; i < 20; i++ {
		f.Factorial(i)
		f.Print(i)
	}
}
