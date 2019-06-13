package queue

import "fmt"

//根据链表实现队列效果
//定义一个链表的节点
type ListNode struct {
	val  interface{}
	next *ListNode
}

//定义一个链表队列
type LinkedListQueue struct {
	head   *ListNode //头节点
	tail   *ListNode //尾部节点
	length int       //队列长度
}

//新建一个链表队列
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

//向链表队列追加元素
func (this *LinkedListQueue) EnQueue(v interface{}) {
	//创建一个写入节点
	node := &ListNode{v, nil}
	//当链表队列为空时
	if this.tail == nil {
		this.tail = node
		this.head = node
	} else {
		//当链表队列为非空时
		//更新当前链表队列最后一个元素为node
		this.tail.next = node
		//node 为最后一个元素
		this.tail = node
	}
}

//向链表队列中清除头部元素 并返回头部元素
func (this *LinkedListQueue) DeQueue() interface{} {
	if this.tail == nil {
		return nil
	}
	v := this.head.val
	//移除链表队列v元素 需要把链表头部位置后移
	this.head = this.head.next
	//维护链表的长度
	this.length--
	return v
}

//输出链表队列信息数据
func (this *LinkedListQueue) String() interface{} {
	if this.tail == nil {
		fmt.Println("Queue is empty")
	}
	//拼接队列头部
	result := "head"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%v", cur.val)
	}
	//拼接队列尾部
	result += "<-tail"
	return result
}
