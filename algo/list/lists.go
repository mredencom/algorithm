package list

import (
	"algorithm/algo/containers"
	"algorithm/algo/utils"
)

type List interface {
	// 获取指定索引的
	Get(index int) (interface{}, bool)
	// 删除指定索引位置元素
	Remove(index int)
	// 添加一个元素
	Add(values ...interface{})
	// 建立一个容器
	Contains(values ...interface{}) bool
	// 比较下
	Sort(comparator utils.Comparator)
	// 交换两个值位置
	Swap(index1, index2 int)
	// 多个插入
	Insert(index int, values ...interface{})
	// 单个插入接口
	Set(index int, value interface{})

	// 需要实现的接口
	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
