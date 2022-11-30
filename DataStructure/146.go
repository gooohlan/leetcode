package DataStructure

import (
	"container/list"
)

type DLinkedNode struct {
	key, value int
	next, prev *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func (n *DLinkedNode) insert(x *DLinkedNode) {
	x.prev = n.next
	x.next = n
	n.prev.next = x
	n.prev = x
}

func (n *DLinkedNode) remove(x *DLinkedNode) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

type LRUCache struct {
	size, capacity int
	head, tail     *DLinkedNode
	cacheMap       map[int]*DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		cacheMap: make(map[int]*DLinkedNode),
	}

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cacheMap[key]; !ok {
		return -1
	} else {
		l.moveHead(node)
		return node.value
	}
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cacheMap[key]; !ok {
		node = initDLinkedNode(key, value)
		l.cacheMap[key] = node
		l.addHead(node)
		l.size++
		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cacheMap, removed.key)
			l.size--
		}
	} else {
		node.value = value
		l.moveHead(node)
	}
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (l *LRUCache) moveHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addHead(node)
}

func (l *LRUCache) addHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

type LRUCache2 struct {
	size     int
	cacheMap map[int]*list.Element
	values   *list.List
}

type Node struct {
	key, val int
}

func Constructor2(capacity int) *LRUCache2 {
	return &LRUCache2{
		size:     capacity,
		cacheMap: make(map[int]*list.Element),
		values:   list.New(),
	}
}

func (l *LRUCache2) Get(key int) int {
	if node, ok := l.cacheMap[key]; !ok {
		return -1
	} else {
		l.values.MoveToFront(node)
		return node.Value.(Node).val
	}
}

func (l *LRUCache2) Put(key int, value int) {
	if node, ok := l.cacheMap[key]; !ok {
		node = l.values.PushFront(Node{key, value})
		l.cacheMap[key] = node
		if l.values.Len() > l.size {
			back := l.values.Back()
			l.values.Remove(back)
			delete(l.cacheMap, back.Value.(Node).key)
		}
	} else {
		l.values.MoveToFront(node)
		data := node.Value.(Node)
		data.val = value
		node.Value = data
	}
}
