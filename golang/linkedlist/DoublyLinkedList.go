package linkedlist

// 双向链表
// 创建一个链表的节点
type Node struct {
	data int
	next *Node //后置节点
	prev *Node // 前置节点
}

// 双向链表的结构
type LinkList struct {
	head *Node // 头节点
	tail *Node // 尾部节点
}

// 插入头部元素
func (list *LinkList) InsertFirst(i int) {
	data := &Node{data: i}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

// 插入尾部元素
func (list *LinkList) InsertLast(i int) {
	data := &Node{data: i}
	// 1.当头结点为空的时候
	if list.head == nil {
		list.head = data
		list.tail = data
		return
	}
	// 2.当尾部节点不为空的时候
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.head
	}
	list.tail = data
}

// 删除指定值节点
func (list *LinkList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == i {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	// 当下一个节点不为空的时候
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// 删除指定索引位置的值
func (list *LinkList) RemoveByIndex(i int) bool {
	// 查看链表是否存在 不存在直接返回
	if list.head == nil {
		return false
	}
	// 索引i小于0是非法索引
	if i < 0 {
		return false
	}
	// 当头结点为0 需要把头结点指向下一个节点 换掉头结点就是删除
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return true
	}
	// 3.第三种情况
	current := list.head
	// 循环到指定索引i的位置
	for u := i; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

// 查找指定的值、=
func (list *LinkList) SearchValue(i int) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}

// 获取第一个元素
func (list *LinkList) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

// 获取末尾元素
func (list *LinkList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	// 把当前节点 循环到最后一个 然后就数据返回回去
	current := list.head
	for current != nil {
		current = current.next
	}
	return current.next.data, true
}

// 获取双向链表的长度
func (list *LinkList) GetSize() int {
	count := 0
	current := list.head
	for current.next != nil {
		count += 1
		current = current.next
	}
	return count
}

// 从链表头开始取出数据
func (list *LinkList) GetItemsFromHead() (items []int) {
	current := list.head
	for current != nil {
		items = append(items, current.next.data)
		current = current.next
	}
	return
}

// 从链表尾开始取出数据
func (list *LinkList) GetItemsFromTail() (items []int) {
	current := list.tail
	for current != nil {
		items = append(items, current.prev.data)
		current = current.prev
	}
	return
}
