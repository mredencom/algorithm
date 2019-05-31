package linkedlist

import "fmt"

//单向链表
//定义一个链表的节点结构
type ListNode struct {
	next  *ListNode
	value interface{}
}

//构造一个链表的结构
type LinkedList struct {
	head   *ListNode
	length uint
}

//创建一个节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

//获取下一个元素
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

//去的当前的节点值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

//建立一个链表结构
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在末一个节点后面插入一个节点 返回bool
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	//判断该节点是否存在
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	//把旧的节点移到新的节点下一个元素
	newNode.next = oldNext
	//增加节点长度
	this.length++
	return true
}

//插入当前节点的前置节点 返回bool
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

//打印方法 打印链表使用的
func (this *LinkedList) Print() {
	//取头节点
	cur := this.head.next
	format := ""
	//判断当前节点下一个节点是否为nil
	for nil != cur {
		format += fmt.Sprintf("%v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "<-"
		}
	}
	fmt.Println(format)
}

//在链表头部插入节点 返回 bool
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertBefore(this.head, v)
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	//循环找到链表的尾部
	for cur.next != nil {
		cur = cur.next
	}
	//最后插入尾部
	return this.InsertAfter(cur, v)
}

//给出指定索引查找对应节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	//取的第一个节点
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除的传入的节点 返回bool
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	//判断节点是否存在
	if p == nil {
		return false
	}
	//取当前节点
	cur := this.head.next
	//头节点  就是我们叫的前置节点
	pre := this.head
	//循环链表
	for cur != nil {
		//找到了删除节点
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}
	//当前节点==nil 就没有找到要删除的节点
	if cur == nil {
		return false
	}
	//做删除操作  就是把前置节点下一个节点 替换掉
	pre.next = p.next
	//然后释放掉删除的几点
	p = nil
	//减少链表的长度
	this.length--
	return true
}
