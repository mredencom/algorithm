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

//测试链表是否有环
//go test -v -run TestLinkedList_HasCycle -o linkedListAlgo_test.go
func TestLinkedList_HasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle())
}

//测试两个有序的链表
//go test -v -run TestLinkedList_MergeSortedList -o linkedListAlgo_test.go
func TestLinkedList_MergeSortedList(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{value: 7, next: n5}
	n3 := &ListNode{value: 5, next: n4}
	n2 := &ListNode{value: 3, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l1 := &LinkedList{head: &ListNode{next: n1}}

	n10 := &ListNode{value: 10}
	n9 := &ListNode{value: 8, next: n10}
	n8 := &ListNode{value: 6, next: n9}
	n7 := &ListNode{value: 4, next: n8}
	n6 := &ListNode{value: 2, next: n7}
	l2 := &LinkedList{head: &ListNode{next: n6}}
	t.Log("链表L1:")
	l1.Print()
	t.Log("链表L2:")
	l2.Print()
	t.Log("合并链表L1和L2:")
	MergeSortedList(l1, l2).Print()

}

//测试两个有序的链表
//go test -v -run TestLinkedList_MergeSortedList2 -o linkedListAlgo_test.go
func TestLinkedList_MergeSortedList2(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{value: 7, next: n5}
	n3 := &ListNode{value: 5, next: n4}
	n2 := &ListNode{value: 3, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l1 := &LinkedList{head: &ListNode{next: n1}}

	n10 := &ListNode{value: 10}
	n9 := &ListNode{value: 8, next: n10}
	n8 := &ListNode{value: 6, next: n9}
	n7 := &ListNode{value: 4, next: n8}
	n6 := &ListNode{value: 2, next: n7}
	l2 := &LinkedList{head: &ListNode{next: n6}}
	t.Log("链表L1:")
	l1.Print()
	t.Log("链表L2:")
	l2.Print()
	t.Log("合并链表L1和L2:")
	//MergeSortedList2(l1, l2).Print()

}
//删除倒数第N个节点
//go test -v -run TestLinkedList_DeleteBottomN -o linkedListAlgo_test.go
func TestLinkedList_DeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}


