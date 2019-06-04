package stack

import "fmt"

//基于数组使用栈的操作
type ArrayStack struct {
	//数据源
	data []interface{}
	//栈顶标识(指针)
	top int
}

//创建一个栈 基于同一个栈操作 需要使用指针
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

//根据栈顶指针可以判断这个栈是否为空栈 空返回true
func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

//向栈添加元素
func (this *ArrayStack) Push(v interface{}) {
	//增加栈的指针位置
	if this.top < 0 {
		this.top = 0
	} else {
		//这个是在原有的基础上增加一个元素
		this.top += 1
	}
	if this.top > len(this.data)-1 {
		//当栈顶元素长度长于定义数组长度 需要使用slice向后增加 增加定义栈的内存空间
		this.data = append(this.data, v)
	} else {
		//向栈添加到指定指针位置的元素
		this.data[this.top] = v
	}
}

//删除栈顶的元素 并返回其元素
func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.top]
	this.top -= 1
	return v
}

//取的栈顶的元素
func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

//清除栈的元素
func (this *ArrayStack) Flush() {
	this.top = -1
}

//打印出栈的内容
func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
