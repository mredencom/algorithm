package stack

//定义栈需要使用的方法
type Stack interface {
	//入栈的方法
	Push(v interface{})
	//删除元素并返回元素
	Pop() interface{}
	//判断栈是否为空
	IsEmpty() bool
	//栈顶元素
	Top() interface{}
	//清空栈-元素
	Flush()
}
