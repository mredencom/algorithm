package bubblesort

//冒泡排序 插入排序  选择排序
//1.冒泡排序 a 是待排序的数组 n 是数组的长度 属于稳定排序 时间复杂度O(n2)
func BubbleSort(a []int, n int) {
	//数组长度小于2就不需要排序
	if n < 2 {
		return
	}
	//定义一个初始的状态
	flag := true
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] < a[j+1] {
				//交换数据
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}

}

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
