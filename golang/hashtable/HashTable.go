package hashtable

import "hash/fnv"

type TableItem struct {
	key  string
	data int
	next *TableItem
}

// 定义一个hashtable
type HashTable struct {
	data [256]*TableItem
}

// 增加一个像hash table 增加一个值
func (table *HashTable) Add(key string, i int) {
	// 构建一个hash 值
	position := GenerateHash(key)
	current := table.data[position]
	// 当前节点没有值 直接可以赋值
	if current == nil {
		// 添加元素
		table.data[position] = &TableItem{data: i, key: key}
	}
	// 当前节点有值 就需要考虑next节点
	for current.next != nil {
		current = current.next
	}
	current.next = &TableItem{key: key, data: i}
}

// 获取一个值
func (table *HashTable) Get(key string) (int, bool) {
	position := GenerateHash(key)
	current := table.data[position]
	// 循环找出和key 相同的的节点
	for current != nil {
		if current.key == key {
			return current.data, true
		}
		current = current.next
	}
	return 0, false
}

// 塞入一个值 set
func (table *HashTable) Set(key string, value int) bool {
	position := GenerateHash(key)
	current := table.data[position]
	for current != nil {
		if current.key == key {
			current.data = value
			return true
		}
		current = current.next
	}
	return false
}

// 删除一个值
func (table *HashTable) Remove(key string) bool {
	position := GenerateHash(key)
	//  当数据不存在的时候 判断下返回false
	if table.data[position] == nil {
		return false
	}
	// 当前key 就是当前节点时候
	if table.data[position].key == key {
		// 把下一个节点
		table.data[position] = table.data[position].next
		return true
	}
	current := table.data[position]
	for current != nil {
		if current.next.key == key {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// 创建一个hash 值
func GenerateHash(key string) uint8 {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(key))
	return uint8(hash.Sum32() % 256)
}
