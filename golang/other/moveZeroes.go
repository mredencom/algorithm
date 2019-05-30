package main

import "fmt"

//移动0算法
//方法1
func moveZeroes(nums []int) []int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			if count > 0 {
				nums[i-count], nums[i] = nums[i], nums[i-count]
			}
		}
	}
	return nums
}

//方法2
func moveZeroes2(nums []int) []int {
	count := 0
	for i := 0; i < len(nums); i++ {
		//找出不为0的个数
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	//把不为0后面重置为0
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
func main() {
	a := []int{0, 1, 0, 3, 12}
	z1 := moveZeroes(a)
	z2 := moveZeroes2(a)
	fmt.Println(z1, z2)

}
