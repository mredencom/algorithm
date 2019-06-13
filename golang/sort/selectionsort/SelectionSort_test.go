package selectionsort

import "testing"

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
