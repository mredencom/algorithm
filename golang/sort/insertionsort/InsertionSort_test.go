package insertionsort

import "testing"

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