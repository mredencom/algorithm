package binarysearch

//二分查找 每次都是折半进行查找 O(log n)

//参数说明 a 是代查询的的有序数组  v是指定的值
func BinarySearch(a []int, v int) int {
	//取出数组的长度
	n := len(a)
	if n == 0 {
		return -1
	}
	low, high := 0, n-1
	for low <= high {
		//取出中间值
		mid := (low + high) / 2
		if mid == v {
			//找到就返回其中的值
			return mid
		} else if mid < v {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	//没有查找到就返回-1标识没有找到
	return -1
}

//使用递归的方法进行二分查找 待查找数组a 待查找值v
func BinarySearchRecursive(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}
	return bs(a, v, 0, n-1)
}

//待查找的数组 a 待查找的值v  数组的开始的下标low 和结束下标
func bs(a []int, v int, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if a[mid] == v {
		return v
	} else if a[mid] > v {
		return bs(a, v, low, mid-1)
	} else {
		return bs(a, v, mid+1, high)
	}
}
