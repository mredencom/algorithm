package linkedlist

import "testing"

//go test -v -run TestLinkedList_InsertAfter -o singlelinkedlist_test.go
//结果:9->8->7->6->5->4->3->2->1
func TestLinkedList_InsertAfter(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i < 10; i++ {
		l.InsertAfter(l.head, i)
	}
	l.Print()
}

//go test -v -run TestLinkedList_InsertBefore -o singlelinkedlist_test.go
//结果:9999->9->8->7->6->5->4->3->2->1
func TestLinkedList_InsertBefore(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i < 10; i++ {
		l.InsertAfter(l.head, i)
	}
	l.InsertBefore(l.head.next, 9999)
	l.InsertBefore(l.head.next, 8888)
	//t.Log(l.head, l.head.next)
	l.Print()
}

//go test -v -run TestLinkedList_FindByIndex -o singlelinkedlist_test.go
func TestLinkedList_FindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i < 10; i++ {
		l.InsertAfter(l.head, i)
	}
	l.InsertBefore(l.head.next, 9999)
	l.InsertBefore(l.head.next, 8888)
	t.Log(l.head.next)
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(20))
}

//go test -v -run TestLinkedList_DeleteNode -o singlelinkedlist_test.go
//删除节点测试
func TestLinkedList_DeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i < 3; i++ {
		l.InsertToTail(i)
	}
	//初始打印
	l.Print()
	//第一次删除打印
	t.Log(l.DeleteNode(l.head.next))
	l.Print()
	//第二次删除打印 (这个地方应该是删除失败的.经过第一次删除,第二次删除已经没有第二个节点了.所以会删除失败的)
	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

//go test -v -run TestLinkedList_InsertToTail -o singlelinkedlist_test.go
//测试插入节点尾部
//结果:1<-2<-3<-4<-5<-6<-7<-8<-9<-10<-11<-12<-13<-14<-15<-16<-17<-18<-19<-20<-21<-22<-23<-24<-25<-26<-27<-28<-29
func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i < 30; i++ {
		l.InsertToTail(i)
	}
	l.Print()
}
