package algorithm

import "fmt"

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
func (this *LinkedList) Reverse()  {
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
