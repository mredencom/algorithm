package algorithm

import (
	"fmt"
)

//单链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

//单链表
type LinkedList struct {
	head *ListNode
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%v", cur.value)
		cur = cur.next
		if nil != cur {
			format += "<-"
		}
	}
	fmt.Println(format)
}

//单链表反转
//时间复杂度O(n)
func (this *LinkedList) Reverse() {
	//判断链表的长度是否有两个节点 没有两个节点就不用反转了
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}
	var pre *ListNode = nil
	cur := this.head.next
	for cur != nil {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}
	this.head.next = pre
}

//判断链表是否有环
//这里使用的快慢指针 当快慢指针相遇就说明有环
func (this *LinkedList) HasCycle() bool {
	if this.head != nil {
		slow := this.head
		fast := this.head
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

//合并两个有序的
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	//判断了l1,l2两个链表
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}
	//创建一个新的链表 保存合并之后的链表
	l := &LinkedList{&ListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for curl1 != nil && curl2 != nil {
		if curl1.value.(int) > curl2.value.(int) {
			//把较小的放在链表的前面
			cur.next = curl2
			//把链表的下一个赋值curl2
			curl2 = curl2.next
		} else {
			//把较小的放在链表的前面
			cur.next = curl1
			//把链表的下一个赋值curl1
			curl1 = curl1.next
		}
		//最后挂载链表的头部
		cur = cur.next
	}
	//if nil != curl1 {
	//	cur.next = curl1
	//} else if curl2 != nil {
	//	cur.next = curl2
	//}
	return l
}

//func MergeSortedList2(l1, l2 *LinkedList) *ListNode  {
//	//判断了l1,l2两个链表
//	if l1 == nil {
//		return l2.head
//	}
//	if l2 == nil {
//		return l1.head
//	}
//	//创建一个新的链表 保存合并之后的链表 需要排好序
//	//var l *ListNode
//	var l *LinkedList
//	if l1.head.value.(int) < l2.head.value.(int) {
//		l = l1
//		l.head.next = MergeSortedList2(l1, l2)
//	} else {
//		l = l2
//		l.head.next = MergeSortedList2(l1, l2)
//	}
//	return l
//}

//删除倒数第N个节点
func (this *LinkedList) DeleteBottomN(n int) {
	slow, fast := this.head, this.head
	if n <= 0 || fast == nil || fast.next == nil {
		return
	}
	//循环找到第n个点
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return
	}
	fmt.Println(slow)
	//利用fast找到n点
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}
