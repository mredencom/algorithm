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




