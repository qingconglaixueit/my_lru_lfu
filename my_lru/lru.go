package my_lru

// hashmap + 双向链表
import "fmt"
import "my_lru_lfu/double_link"

type LRUCache struct {
	Capacity int
	find     map[interface{}]*double_link.Node
	list     *double_link.List
	k        int
	Count    int
}

func InitLRU(Capacity int) *LRUCache {
	return &LRUCache{
		Capacity: Capacity,
		list:     double_link.InitList(Capacity),
		find:     make(map[interface{}]*double_link.Node,  Capacity),
	}
}

func (l *LRUCache) Get(Key interface{}) interface{} {
	if Value, ok := l.find[Key]; ok {
		node := Value
		l.list.Remove(node)
		l.list.AppendToHead(node)
		return node.Value
	} else {
		node := l.list.RemoveTail()
		fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "要开始删除节点了哦", 0x1B)
		delete(l.find, node.Key)
		node.Key = Key
		l.find[Key] = node
		l.list.AppendToHead(node)
		l.Count++
		return -1
	}
}

func (l *LRUCache) Put(Key, Value interface{}) {
	if v, ok := l.find[Key]; ok {
		node := v
		l.list.Remove(node)
		node.Value = Value
		l.list.AppendToHead(node)
	} else {
		node := double_link.InitNode(Key, Value)
		// 缓存已经满了
		if l.list.Size >= l.list.Capacity {
			oldNode := l.list.Remove(nil)
			delete(l.find, oldNode.Value)
		}
		l.list.AppendToHead(node)
		l.find[Key] = node
	}
}

func (l *LRUCache) String() string {
	return l.list.String()
}
