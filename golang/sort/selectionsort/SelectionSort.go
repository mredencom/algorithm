package selectionsort

//选择排序 a 是待排序的数组 n 是数组的长度  这个是破坏原有的顺序  O(n2)
func SelectionSort(a []int, n int) {
	//数组长度小于2就不需要排序
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		//首先选择一个最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				//寻找最小的下标索引
				minIndex = j
			}
		}
		//然后给当前给最小的进行交换
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}
