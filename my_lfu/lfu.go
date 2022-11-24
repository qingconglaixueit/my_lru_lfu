package my_lfu

import (
	"fmt"
	"math"
	"my_lru_lfu/double_link"
	"strings"
)

// 两个 hashmap + 双向链表
// 一个 hashmap 存放 key 对应的节点
// 一个 hashmap 存放 频次 对应的链表
type LFUNode struct {
	freq int
	node *double_link.Node
}

func InitLFUNode(Key, Value interface{}) *LFUNode {
	return &LFUNode{
		freq: 0,
		node: double_link.InitNode(Key, Value),
	}
}

type LFUCache struct {
	capacity int
	find     map[interface{}]*LFUNode
	freq_map map[int]*double_link.List
	Size     int
	Count    int
}

func InitLFUCahe(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		find:     map[interface{}]*LFUNode{},
		freq_map: map[int]*double_link.List{},
	}
}

// 更新节点的频率
func (l *LFUCache) updateFreq(node *LFUNode) {
	freq := node.freq
	// 删除
	node.node = l.freq_map[freq].Remove(node.node)
	if l.freq_map[freq].Size == 0 {
		delete(l.freq_map, freq)
	}

	freq++
	node.freq = freq
	if _, ok := l.freq_map[freq]; !ok {
		l.freq_map[freq] = double_link.InitList(10)
	}
	l.freq_map[freq].Append(node.node)
}
func findMinNum(fmp map[int]*double_link.List) int {
	min := math.MaxInt32
	for Key, _ := range fmp {
		min = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}(min, Key)
	}
	return min
}
func (l *LFUCache) Get(Key interface{}) interface{} {
	if _, ok := l.find[Key]; !ok {
		min_freq := findMinNum(l.freq_map)
		list := l.freq_map[min_freq]
		node := list.Head
		fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "要开始删除节点了哦", 0x1B)
		// 先取到这个节点的地址
		newNode := l.find[node.Key]
		// 从节点的映射中删掉这个节点
		delete(l.find, node.Key)
		// 赋值新的key
		newNode.node.Key = Key
		l.find[Key] = newNode
		// 在链表中真正的删除这个节点,并且删除后如果链表的长度为0,在频率映射表中吧这个链表删掉
		list.Remove(newNode.node)
		if list.Size == 0 {
			delete(l.freq_map, newNode.freq)
		}
		newNode.freq = 0
		if _, ok := l.freq_map[0]; !ok {
			l.freq_map[0] = double_link.InitList(10)
		}
		l.freq_map[0].Append(newNode.node)
		l.updateFreq(newNode)
		l.Count++
		return -1
	}
	node := l.find[Key]
	l.updateFreq(node)
	return node.node.Value
}

func (l *LFUCache) Put(Key, Value interface{}) {
	if l.capacity == 0 {
		return
	}
	// 命中缓存
	if _, ok := l.find[Key]; ok {
		node := l.find[Key]
		node.node.Value = Value
		l.updateFreq(node)
	} else {
		if l.capacity == l.Size {
			// 找到一个最小的频率
			min_freq := findMinNum(l.freq_map)
			node := l.freq_map[min_freq].Pop()
			lfuNode := &LFUNode{
				node: node,
				freq: 1,
			}
			l.find[Key] = lfuNode
			delete(l.find, node.Key)
			l.Size--
		}
		node := InitLFUNode(Key, Value)
		node.freq = 1
		l.find[Key] = node
		if _, ok := l.freq_map[node.freq]; !ok {
			l.freq_map[node.freq] = double_link.InitList(math.MaxInt32)
		}
		node.node = l.freq_map[node.freq].Append(node.node)
		l.Size++
	}
}

func (l *LFUCache) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, "\n----------------------------------------")
	for k, v := range l.freq_map {
		fmt.Fprintf(&builder, "Freq = %d : ", k)
		fmt.Fprintln(&builder, v.String())
	}
	fmt.Fprintln(&builder, "----------------------------------------")
	return builder.String()
}
