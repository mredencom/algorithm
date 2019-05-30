package array

import (
	"errors"
	"fmt"
)

/**
 * 这里是学习数组的操作和使用(包含数组的插入 删除 ) 这里定义数组都是int类型
 * @author jt 2019-5-30 10:20:12
 */

//定义一个结构体
type Array struct {
	data   []int //数组
	length int   //长度
}

// 我们这里做一个建立数组的方法
func NewArray(capacity uint) *Array {
	//因为我们数组的长度不能为0 因为0初始化没有意义
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

//取当前数组的长度
func (this *Array) Len() uint {
	return uint(this.length)
}

//判断数组是否越界  true:越界  false:没有越界 n~n-1
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

//这个根据索引查找数组中的数据
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("数组下标越界")
	}
	return this.data[index], nil
}

//给指定的值插入到指定索引上. 返回错误
func (this *Array) Insert(index uint, v int) error {
	//第一步检查数组的时候已经满载
	if this.Len() == uint(cap(this.data)) {
		return errors.New("数组已经满了!不可增加值")
	}
	if this.isIndexOutOfRange(index) && index != uint(this.length) {
		return errors.New("数组下标越界!")
	}
	for i := this.length; int(index) < i; i-- {
		//index后的元素后移
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

//追加元素到数组的尾部
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

//删除指定索引上的值
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("索引下标越界")
	}
	//去除指定索引上的值
	v := this.data[index]
	//把删除后的元素的其他索引元素前移
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	//更新数组长度
	this.length--
	return v, nil
}

//输出一个数组
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		//拼接数组
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
