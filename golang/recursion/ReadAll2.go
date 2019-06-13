package recursion

//全排列 [1,2,3]
func Permute(nums []int) [][]int {
	l := len(nums)
	result := make([][]int, 0)
	if l <= 1 {
		return append(result, nums)
	}

}
