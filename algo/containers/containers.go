package containers

// 建立容器的操作接口
type Container interface {
	Empty() bool
	Clear()
	Size() int
	Values() []interface{}
}
