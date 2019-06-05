package queue

import "fmt"

//@todo 这里使用数组实现队列效果
//定义一个队列数据结构
type ArrayQueue struct {
	q        []interface{} //值
	capacity int           //容量
	head     int           //头节点
	tail     int           //尾部节点
}

//新建一个数组队列
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

//检查队列是否还可以追加元素 可以追加就加在队列上 返回状态
func (this *ArrayQueue) EnQueue(v interface{}) bool {
	//检查队列长度是否等于tail节点
	if this.capacity == this.tail {
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

//取出头节点数据 并删除
func (this *ArrayQueue) DeQueue() interface{} {
	//检查对了是否有数据
	if this.head == this.tail {
		return nil
	}
	v := this.q[this.head]
	//头节点后移
	this.head++
	return v
}

//输出队列数据
func (this *ArrayQueue) String() string {
	//检查队列是否有数据
	if this.head == this.tail {
		return "Queue is empty"
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%v", this.q[i])
	}
	result += "<-tail"
	return result
}
