package double_link

import (
	"fmt"
	"strings"
)

// 节点数据结构
type Node struct {
	Key        interface{}
	Value      interface{}
	Prev, Next *Node
}

// 使用 strings.Builder 的方式处理字符串，效率更高
func (n Node) String() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "{%v : %v}", n.Key, n.Value)
	return builder.String()
}

func InitNode(Key, Value interface{}) *Node {
	return &Node{
		Key:   Key,
		Value: Value,
	}
}

// 链表数据结构
type List struct {
	Capacity int
	Head     *Node
	Tail     *Node
	Size     int
}

// 列表的初始化方法
func InitList(capcity int) *List {
	return &List{
		Capacity: capcity,
		Size:     0,
	}
}

// 头插法
func (l *List) addHead(node *Node) *Node {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		l.Head.Prev = nil
		l.Tail.Next = nil
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
		l.Head.Prev = nil
	}
	l.Size++
	return node
}

// 尾插法
func (l *List) addTail(node *Node) *Node {
	if l.Tail == nil {
		l.Tail = node
		l.Head = node
		l.Head.Prev = nil
		l.Tail.Next = nil
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = node
		l.Tail.Next = nil
	}
	l.Size++
	return node
}

// 删除某一个双向链表中的节点
func (l *List) remove(node *Node) *Node {
	// 如果node==nil,默认删除尾节点
	if node == nil {
		node = l.Tail
	}
	if node == l.Tail {
		l.RemoveTail()
	} else if node == l.Head {
		l.RemoveHead()
	} else {
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
		l.Size--
	}
	return node
}

// 弹出头结点
func (l *List) Pop() *Node {
	return l.RemoveHead()
}

// 添加节点,默认添加到尾部
func (l *List) Append(node *Node) *Node {
	return l.addTail(node)
}
func (l *List) AppendToHead(node *Node) *Node {
	return l.addHead(node)
}

// 删除尾节点
func (l *List) RemoveTail() *Node {
	if l.Tail == nil {
		return nil
	}
	node := l.Tail
	if node.Prev != nil {
		l.Tail = node.Prev
		l.Tail.Next = nil
	} else {
		l.Tail = nil
		l.Head = nil
	}
	l.Size--
	return node
}

// 删除头结点
func (l *List) RemoveHead() *Node {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	if node.Next != nil {
		l.Head = node.Next
		l.Head.Prev = nil
	} else {
		l.Tail = nil
		l.Head = nil
	}
	l.Size--
	return node
}

func (l *List) Remove(node *Node) *Node {
	return l.remove(node)
}

func (l *List) String() string {
	p := l.Head
	builder := strings.Builder{}
	for p != nil {
		fmt.Fprintf(&builder, "[%d, %d]", p.Key,p.Value)
		p = p.Next
		if p != nil {
			fmt.Fprintf(&builder, " => ")
		}
	}
	return builder.String()
}
