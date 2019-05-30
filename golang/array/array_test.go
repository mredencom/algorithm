package array

import "testing"

//这里是一个测试array 文件的测试方法集

//测试数组插入方法 go test -v -run TestArray_Insert -o array_test.go
func TestArray_Insert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		//循环入数组
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	_ = arr.Insert(uint(6), 999)
	arr.Print()
	_ = arr.InsertToTail(555)
	arr.Print()
}

//测试删除方法 go test -v -run TestArray_Delete -o array_test.go
func TestArray_Delete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	//这里做循环删除
	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
	arr.Print()
}

//测试 go test -v -run TestArray_Find -o array_test.go
func TestArray_Find(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
}
