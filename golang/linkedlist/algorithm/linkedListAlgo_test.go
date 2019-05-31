package algorithm

import "testing"

var l *LinkedList

//测试 初始化数据
func init() {
	n5 := &ListNode{value: 5}
	n4 := &ListNode{value: 4, next: n5}
	n3 := &ListNode{value: 3, next: n4}
	n2 := &ListNode{value: 2, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l = &LinkedList{head: &ListNode{next: n1}}
}

//go test -v -run TestLinkedList_Reverse -o linkedListAlgo_test.go
func TestLinkedList_Reverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}
