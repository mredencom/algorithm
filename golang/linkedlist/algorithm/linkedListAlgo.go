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
func (this *ListNode) Reverse() {

}
