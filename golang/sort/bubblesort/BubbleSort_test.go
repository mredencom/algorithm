package bubblesort

import "testing"

//测试冒泡排序
func TestBubbleSort(t *testing.T) {
	a := []int{95, 45, 15, 78, 84, 51, 24, 12}
	t.Log("排序前:", a)
	BubbleSort(a, len(a))
	t.Log("排序后:", a)
}

//并发压测
func BenchmarkBubbleSort(b *testing.B) {
	//b.StartTimer()
	a := []int{95, 45, 15, 78, 84, 51, 24, 12}
	b.Log("排序前:", a)
	BubbleSort(a, len(a))
	b.Log("排序后:", a)
	//b.StopTimer()
}
func ExampleBubbleSort() {
}

