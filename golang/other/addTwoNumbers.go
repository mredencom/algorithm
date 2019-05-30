package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cursor := res
	carry := 0
	for {
		val := l1.Val + l2.Val + carry
		carry = val / 10
		val = val % 10
		cursor.Val = val

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			l1 = &ListNode{}
			l1.Val = 0
			l1.Next = nil
		}
		if l2 == nil {
			l2 = &ListNode{}
			l2.Val = 0
			l2.Next = nil
		}

		cursor.Next = &ListNode{}
		cursor = cursor.Next
	}
	if carry != 0 {
		cursor.Next = &ListNode{Val: carry}
	}
	return res
}

func Example() {
	//测试创建一个链表
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushBack(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//fmt.Println(l)
}

func main() {
	//l1 := &ListNode{1, 2, 3, &ListNode{}}
	//l2 := &ListNode{4, 5, 6, &ListNode{}}
	l1 := new(ListNode)
	l1.Val = 1
	//l1.Next.Val = 2
	l2 := new(ListNode)
	l2.Val = 5
	//l2.Next.Val = 4
	l1.Next = l2
	//a := addTwoNumbers(l1, l2)
	//fmt.Println(a)
	//Example()

	Shownode(l1)
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
