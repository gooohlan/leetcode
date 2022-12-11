package WeeklyContest

type Allocator struct {
	list []int
	len  int
}

func Constructor(n int) Allocator {
	return Allocator{
		list: make([]int, n),
		len:  n,
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	len := 0
	i := 0
	for ; i < this.len; i++ {
		if this.list[i] == 0 {
			len++
		} else {
			len = 0
		}
		if len == size {
			break
		}
	}
	if len == 0 || len < size {
		return -1
	}
	i = i - len + 1
	for j := 0; j < size; j++ {
		this.list[i+j] = mID
	}
	return i
}

func (this *Allocator) Free(mID int) int {
	count := 0
	for i, v := range this.list {
		if v == mID {
			count++
			this.list[i] = 0
		}
	}
	return count
}
