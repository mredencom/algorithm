package stack

import "fmt"

type Browser struct {
	forwardStack Stack
	backStack    Stack
}

//建立一个Browser
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

//判断是否为空
func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}
func (this *Browser) CanBack() bool {
	if this.backStack.IsEmpty() {
		return false
	}
	return true
}

//打开地址
func (this *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	this.forwardStack.Flush()
}

//增加一个
func (this *Browser) PushBack(addr string) {
	this.backStack.Push(addr)
}

//取最前面一个
func (this *Browser) Forward() {
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

//删除一个并返回
func (this *Browser) Back() {
	if this.backStack.IsEmpty() {
		return
	}
	top := this.backStack.Pop()
	this.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
