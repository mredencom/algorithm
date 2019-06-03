package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//判断了l1,l2两个链表
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//创建一个新的链表 保存合并之后的链表 需要排好序
	var l *ListNode
	if l1.Val < l2.Val {
		l = l1
		l.Next = mergeTwoLists(l1.Next, l2)
	} else {
		l = l2
		l.Next = mergeTwoLists(l1, l2.Next)
	}
	return l
}

func main() {

	n5 := &ListNode{Val: 9}
	n4 := &ListNode{Val: 7, Next: n5}
	n3 := &ListNode{Val: 5, Next: n4}
	n2 := &ListNode{Val: 3, Next: n3}
	l1 := &ListNode{Val: 1, Next: n2}

	n10 := &ListNode{Val: 10}
	n9 := &ListNode{Val: 8, Next: n10}
	n8 := &ListNode{Val: 6, Next: n9}
	n7 := &ListNode{Val: 4, Next: n8}
	l2 := &ListNode{Val: 2, Next: n7}

	l := mergeTwoLists(l1, l2)
	//fmt.Println(l)
	//Print(l1)
	//Print(l2)
	Print(l)

}

func Print(l *ListNode) {
	cur := l.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%v", cur.Val)
		cur = cur.Next
		if nil != cur {
			format += "<-"
		}
	}
	fmt.Println(format)
}
