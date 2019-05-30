package main

import "fmt"
//leetcode-cn 网站  27.移除元素
func removeElement(nums []int, val int) []int {
	var c []int
	for _, v := range nums {
		if val != v {
			c = append(c, v)
		}
	}
	return c
}
func main() {
	a := []int{3, 2, 2, 3}
	s := removeElement(a, 2)
	fmt.Println(s)
}
