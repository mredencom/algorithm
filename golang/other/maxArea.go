package main

import (
	"fmt"
)

//找出最大的面积区域
func maxArea(height []int) int {
	var max, area int
	start := 0
	end := len(height) - 1 //取数组的长度
	//建立循环取出最大的赋值给max 最后返回max
	for start < end {
		temp := end - start
		if height[start] > height[end] {
			area = temp * height[end]
			end--
		} else {
			area = temp * height[start]
			start++
		}
		if area > max {
			max = area
		}
	}
	return max
}
func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	c := maxArea(arr)
	fmt.Println(c)
}
