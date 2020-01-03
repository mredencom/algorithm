package arraylist

import (
	"algorithm/algo/utils"
	"fmt"
	"strings"
)

// wiki: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
// 为了数组链表空闲过大 or 数组链表过小 在这里我定义两个常量
const (
	growth = float32(2.0)  // 增长一倍
	shrink = float32(0.25) // 当长度为25%就会减小容量
)

// 定义一个数组的链表的结构
type List struct {
	// 链表的数组
	elements []interface{}
	// 数组长度
	size int
}

// 创建一个array list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		// 增加元素
		list.Add(values)
	}
	return list
}

// 向数组链表增加元素
func (list *List) Add(values ...interface{}) {
	// 检查写是否需要增加slice 的长度
	list.growBy(len(values))
	for _, value := range values {
		list.elements = append(list.elements, value)
		list.size++
	}
}

// 根据索引去查找一个元素
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withInRange(index) {
		return nil, false
	}
	return list.elements[index], false
}

// 根据索引删除一个链表元素
func (list *List) Remove(index int) {
	// 判断索引元素是否在在链表上
	if !list.withInRange(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	//list.elements[index:] = append(list.elements[index:], list.elements[index+1:list.size]...)
	list.size--

	// 检查是否需要减少长度
	list.shrink()
}

// 查找元素是否存在于list中 存在返回true 否则 返回 false
// 查找复杂度 n^2
func (list *List) Contains(values ...interface{}) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if searchValue == element {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// 获取链表中所有的值
func (list *List) Values() []interface{} {
	listValues := make([]interface{}, list.size, list.size)
	//listValues = append(listValues, list.elements[:list.size])
	copy(listValues, list.elements[:list.size])
	return listValues
}

// 根据value 找出其索引位置
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

// 判断链表是否为空
func (list *List) Empty() bool {
	return list.size == 0
}

// 获取链表的长度
func (list *List) Size() int {
	return list.size
}

// 清除链表数据
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// 判断索引是否在 链表中
func (list *List) withInRange(index int) bool {
	return index >= 0 && index < list.size
}

// 排序需要实现的方法
func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// 交换数据
func (list *List) Swap(i, j int) {
	if list.withInRange(i) && list.withInRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// 插入数据
func (list *List) Insert(index int, values ...interface{}) {
	if !list.withInRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

// to string
func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
// 增加array list 长度
func (list *List) growBy(n int) {
	// 获取capacity 长度
	currentCap := cap(list.elements)
	if list.size+n >= currentCap {
		newCapacity := int(growth * float32(currentCap+n))
		list.resize(newCapacity)
	}
}

// 减小 array list 长度
func (list *List) shrink() {
	if shrink == 0.0 {
		// shrink 为 0 不做任何操作
		return
	}
	currentCap := cap(list.elements)
	if list.size <= int(float32(currentCap)*shrink) {
		list.resize(list.size)
	}
}

// 重置链表的大小
func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	// 把内容copy 新的slice上
	copy(newElements, list.elements)
	list.elements = newElements
}
