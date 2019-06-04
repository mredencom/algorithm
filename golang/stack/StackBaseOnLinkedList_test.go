package stack

import "testing"

//测试链表栈
//go test -v -run TestNewLinkedListStack -o StackBaseOnLinkedList_test.go
func TestNewLinkedListStack(t *testing.T) {
	s := NewLinkedListStack()
	s.Print()
}

//测试空链表栈
//go test -v -run TestLinkedListStack_IsEmpty -o StackBaseOnLinkedList_test.go
func TestLinkedListStack_IsEmpty(t *testing.T) {
	s := NewLinkedListStack()
	t.Log(s.IsEmpty())
}

//测试push链表栈
//go test -v -run TestLinkedListStack_Push -o StackBaseOnLinkedList_test.go
func TestLinkedListStack_Push(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1111)
	s.Push(222)
	s.Print()
}

//测试Pop链表栈
//go test -v -run TestLinkedListStack_Pop -o StackBaseOnLinkedList_test.go
func TestLinkedListStack_Pop(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1111)
	s.Push(222)
	s.Print()
	t.Log(s.Pop())
	s.Print()
}

//测试Top链表栈
//go test -v -run TestLinkedListStack_Top -o StackBaseOnLinkedList_test.go
func TestLinkedListStack_Top(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1111)
	s.Push(222)
	s.Print()
	t.Log(s.Top())
}

//测试Top链表栈
//go test -v -run TestLinkedListStack_Flush -o StackBaseOnLinkedList_test.go
func TestLinkedListStack_Flush(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1111)
	s.Push(222)
	s.Print()
	s.Flush()
	s.Print()
	t.Log(s.IsEmpty())
}
