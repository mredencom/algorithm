package queue

import "testing"

//测试EnQueue 方法
//go test -v -run TestLinkedListQueue_EnQueue -o QueueBaseOnLinkedList_test.go
func TestLinkedListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()
	//空的数组链表队列
	t.Log(q.String())
	//增加队列元素
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	t.Log(q.String())
}

//测试DeQueue 方法
//go test -v -run TestLinkedListQueue_DeQueue -o QueueBaseOnLinkedList_test.go
func TestLinkedListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()
	//空的数组链表队列
	t.Log(q.String())
	//增加队列元素
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	t.Log(q.String())
	//清除tail一个元素
	t.Log(q.DeQueue())
}
