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

//测试插入排序
func TestInsertionSort(t *testing.T) {
	a := []int{95, 45, 15, 78, 84, 51, 24, 12}
	t.Log("插入排序前:", a)
	InsertionSort(a, len(a))
	t.Log("插入排序后:", a)
}

//并发压测插入排序
func BenchmarkInsertionSort(b *testing.B) {
	a := []int{95, 45, 15, 78, 84, 51, 24, 12}
	b.Log("插入排序前:", a)
	InsertionSort(a, len(a))
	b.Log("插入排序后:", a)
}

//选择排序 测试
func TestSelectionSort(t *testing.T) {
	a := []int{95, 45, 15, 78, 84, 51, 24, 12}
	t.Log("选择排序前:", a)
	SelectionSort(a, len(a))
	t.Log("选择排序后:", a)
}

//选择排序 压力测试
func BenchmarkSelectionSort(b *testing.B) {
	a := []int{95, 45, 15, 78, 84, 51, 24, 12}
	b.Log("选择排序前:", a)
	SelectionSort(a, len(a))
	b.Log("选择排序后:", a)
}
