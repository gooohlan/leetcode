package linked_list

import "container/heap"

// An IntHeap is a min-heap of ints.
type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 使用最小堆存储所有链表头结点节点
// 遍历最小堆,取出最小节点添加到新链表,同时将取出节点下一个节点放入最小堆中
func mergeKLists2(lists []*ListNode) *ListNode {
	h := new(minHeap)
	// 将所有链表的头结点放入最小堆
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	head := &ListNode{}
	pre := head
	for h.Len() > 0 {
		// 取出最小节点
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			// 将最小节点下一节点加入最小堆
			heap.Push(h, node.Next)
		}
		// 最小节点放入链表
		pre.Next = node
		pre = pre.Next
	}
	return head.Next
}
