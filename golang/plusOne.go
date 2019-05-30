package main

import "fmt"
//leetcode-cn.com  66. 加一  算法
/**
		给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
	最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
	你可以假设除了整数 0 之外，这个整数不会以零开头。
	示例 1:
	输入: [1,2,3]
	输出: [1,2,4]
	解释: 输入数组表示数字 123。
	示例 2:

	输入: [4,3,2,1]
	输出: [4,3,2,2]
	解释: 输入数组表示数字 4321。
 */
func plusOne(digits []int) []int {
	//自己循环最后一位
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			//需要最后一位加1
			digits[i]++
			return digits
		} else {
			//这个地方相当于9+1 =10 需要进位
			digits[i] = 0
		}
	}
	//这里返回只处理 99, 999, 9999 这种情况
	return append([]int{1}, digits...)
}

func main() {
	s := []int{6, 2, 9}
	res := plusOne(s)
	fmt.Println(res)
}
