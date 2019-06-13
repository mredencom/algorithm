package insertionsort

//插入排序  a 是待排序的数组 n 是数组的长度 属于稳定排序 时间复杂度O(n2)
func InsertionSort(a []int, n int) {
	//数组长度小于2就不需要排序
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		current := a[i] //循环的当前值
		j := i - 1
		//在做一个循环找到current当前值位置并插入a[i]后面
		for ; j >= 0; j-- {
			if a[j] > current {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		//所以吧当前current值移到当前j的后面
		a[j+1] = current
	}
}