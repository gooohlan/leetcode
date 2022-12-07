package DataStructure

import (
	"container/heap"
	"sort"
)

// 维护两个优先队列,large, small分别记录比中位数大的和比中位数小的
// 当长度为奇数时,保证small比large多一个,此时中位数为small的顶堆
// 长度为偶数时,取两个队列的第一个元素,求平均数
// 当插入队列时,如果数量超出限制(保证small比large多一个),则将队列第一个加入另一个队列即可
type MedianFinder struct {
	large hp // 大顶堆
	small hp // 小顶堆 需要最大的在堆顶,所以取反
}

func Constructor295() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.small.Len() == 0 || num <= -this.small.IntSlice[0] {
		heap.Push(&this.small, -num)
		if this.large.Len()+1 < this.small.Len() { // 小顶堆数量比大顶堆数量多出2,将小顶堆元素移入大顶堆
			heap.Push(&this.large, -heap.Pop(&this.small).(int))
		}
	} else {
		heap.Push(&this.large, num)
		if this.large.Len() > this.small.Len() { // 大顶堆数量大于小顶堆数量,将大顶堆元素移入小顶堆
			heap.Push(&this.small, -heap.Pop(&this.large).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64(-this.small.IntSlice[0])
	}

	return float64(this.large.IntSlice[0]-this.small.IntSlice[0]) / 2
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
