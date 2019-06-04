package stack

import "testing"

//测试新建栈
//go test -v -run TestNewArrayStack -o StackBaseOnArray_test.go
func TestNewArrayStack(t *testing.T) {
	//新建一个栈
	s := NewArrayStack()
	//打印栈
	s.Print()
}

//测试向栈加入元素
//go test -v -run TestArrayStack_Push -o StackBaseOnArray_test.go
func TestArrayStack_Push(t *testing.T) {
	//新建一个栈
	s := NewArrayStack()
	//向栈加入元素
	s.Push(11111)
	s.Push("你好")
	s.Push("hello stack")
	s.Print()
	//取栈顶元素
	t.Log(s.top)
}

//测试栈是否为空
//go test -v -run TestArrayStack_IsEmpty -o StackBaseOnArray_test.go
func TestArrayStack_IsEmpty(t *testing.T) {
	s := NewArrayStack()
	t.Log("空栈: ", s.IsEmpty())
	s.Push(111)
	t.Log("非空栈: ", s.IsEmpty())
}

//测试删除栈顶元素 并返回
//go test -v -run TestArrayStack_Pop -o StackBaseOnArray_test.go
func TestArrayStack_Pop(t *testing.T) {
	//新建一个栈
	s := NewArrayStack()
	s.Push(11111)
	s.Push("你好")
	s.Push("需要删除的栈元素")
	s.Print()
	d := s.Pop()
	t.Log("需要删除的栈元素: ", d)
	t.Log("------------------------")
	s.Print()
}

//测试取栈顶元素
//go test -v -run TestArrayStack_Top -o StackBaseOnArray_test.go
func TestArrayStack_Top(t *testing.T) {
	//新建一个栈
	s := NewArrayStack()
	s.Push(11111)
	s.Push("你好")
	s.Push("待取的栈顶元素")
	s.Print()
	t.Log(s.Top())
}

//测试取栈顶元素
//go test -v -run TestArrayStack_Flush -o StackBaseOnArray_test.go
func TestArrayStack_Flush(t *testing.T) {
	//新建一个栈
	s := NewArrayStack()
	s.Push(11111)
	s.Push("你好")
	s.Print()
	s.Flush()
	s.Print()
}
