package queue

import "testing"

//测试EnQueue 方法
//go test -v -run TestArrayQueue_EnQueue -o QueueBaseOnArray_test.go
func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(10)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	t.Log(q.String())

}

//测试DeQueue 方法
//go test -v -run TestArrayQueue_DeQueue -o QueueBaseOnArray_test.go
func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(10)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	//清除元素
	t.Log(q.String())
	q.DeQueue()
	t.Log(q.String())

}
