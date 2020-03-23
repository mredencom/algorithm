package LinearSearch

// 这里是实现一个线性的检索 返回索引
// todo 这里只是做了一个例子 为了学习使用 可以自由发挥的
func LinearSearch(array []int, number int) int {
	for index, value := range array {
		if number == value {
			return index
		}
	}
	return -1
}
