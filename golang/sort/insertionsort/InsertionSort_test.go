package insertionsort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

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

/// 测试下 InsertionSort2
func TestInsertionSort2(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1 := make([]int, rand.Intn(100-10)+10)
	for i := range array1 {
		array1[i] = random.Intn(100)
	}
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)
	InsertionSort2(array1)
	array2.Sort()
	t.Log("array1:", array1)
	t.Log("array2:", array2)
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}
