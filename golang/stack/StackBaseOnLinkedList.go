package stack

import "fmt"

//使用链表实现栈的结构
//定义一个链表结构
type node struct {
	next *node
	val  interface{}
}

//定义一个栈链表结构
type LinkedListStack struct {
	topNode *node
}

//新建一个链表栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

//判断栈是否为空
func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	}
	return false
}

//向链表栈push加入数据
func (this *LinkedListStack) Push(v interface{}) {
	this.topNode = &node{this.topNode, v}
}

//删除栈顶的元素 并返回栈顶的元素
func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.topNode.val
	//删除元素
	this.topNode = this.topNode.next
	return v
}

//取栈顶表的数据
func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.val
}

//清空链表栈
func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

//打印
func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
	} else {
		cur := this.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
