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

// 插入排序的第二张实现方式 array 这是一个待排序的数组
func InsertionSort2(array []int) {
	// 首先我们要知道 如果是一维数组 长度为2的时候 就根本不需要排序
	if len(array) < 2 {
		return
	}
	for itemIndex, itemValue := range array {
		// 当前面大于元素大于后面一个 就给前面的元素移动到后面上来的
		// 这里的itemIndex大于1的
		for itemIndex != 0 && array[itemIndex-1] > itemValue {
			array[itemIndex] = array[itemIndex-1]
			itemIndex -= 1
		}
		array[itemIndex] = itemValue
	}
}
