package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//定义一个结果数组
	result := make([]int, 0)
	m := make(map[int]int)
	if nums == nil || len(nums) <= 1 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		r, ok := m[target-num]
		if ok {
			//如果成立写入slice
			result = append(result, r, i)
			return result
		} else {
			m[num] = i
			fmt.Println(m)
		}
	}
	return result
}
func main() {
	nums := []int{2, 7, 11, 15}
	t := 9
	sum := twoSum(nums, t)
	fmt.Println(sum)
}
