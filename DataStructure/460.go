package DataStructure

import "container/list"

type Node460 struct {
	key, val, freq int
}

type LFUCache struct {
	minFreq, capacity int                   // 最小访问频率以及容量
	keyTable          map[int]*list.Element // key到val的映射
	freqTable         map[int]*list.List    // 访问频率映射
}

func Constructor460(capacity int) LFUCache {
	l := LFUCache{
		minFreq:   1,
		capacity:  capacity,
		keyTable:  make(map[int]*list.Element, capacity),
		freqTable: make(map[int]*list.List, capacity),
	}
	l.freqTable[l.minFreq] = list.New() // freq为1肯定会用到,提前申请
	return l
}

func (l *LFUCache) Get(key int) int {
	if len(l.keyTable) == 0 {
		return -1
	}
	node, ok := l.keyTable[key]
	if !ok {
		return -1
	}
	l.increaseFreq(node)
	return node.Value.(*Node460).val
}

func (l *LFUCache) Put(key int, value int) {
	if l.capacity <= 0 {
		return
	}
	// 查询key是否存在,存在则直接更新即可
	if node, ok := l.keyTable[key]; ok {
		node.Value.(*Node460).val = value
		l.increaseFreq(node)
		return
	}
	if len(l.keyTable) >= l.capacity { // 缓存已满,移除 freqTable[minFreq] 链表的末尾节点
		back := l.freqTable[l.minFreq].Back()         // 找出最小节点
		l.freqTable[l.minFreq].Remove(back)           // 移除节点
		delete(l.keyTable, back.Value.(*Node460).key) // 从keyTable移除该节点
	}

	// 插入
	node := l.freqTable[1].PushFront(&Node460{key, value, 1})
	l.keyTable[key] = node
	l.minFreq = 1
}

func (l *LFUCache) increaseFreq(node *list.Element) {
	// 取出节点
	data := node.Value.(*Node460)
	l.freqTable[data.freq].Remove(node)                              // 从freq列表中删除节点
	if l.freqTable[data.freq].Len() == 0 && l.minFreq == data.freq { // 最小freqMap列表为空,则更新minFreq
		l.minFreq++
	}
	data.freq++ // 更新节点freq
	if l.freqTable[data.freq] == nil {
		l.freqTable[data.freq] = list.New()
	}
	node = l.freqTable[data.freq].PushFront(data) // 放入freq+1列表中
	l.keyTable[data.key] = node
}
