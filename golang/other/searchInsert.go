package main

import "fmt"

//leetcode-cn 网站 35 搜索插入位置
func searchInsert(nums []int, target int) int {
	var min int
	max := len(nums)
	for min < max {
		mid := min + (max-min)/2
		fmt.Println(mid, min, max)
		if nums[mid] > target {
			max = mid
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}
	}
	return min
}

func main() {
	nums := []int{1, 3, 5, 6, 7}
	res := searchInsert(nums, 5)
	fmt.Println(res)
}
