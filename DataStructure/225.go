package DataStructure

type MyStack struct {
    queue []int
}

func Constructor225() MyStack {
    return MyStack{
        queue: make([]int, 0),
    }
}

func (this *MyStack) Push(x int) {
    this.queue = append([]int{x}, this.queue...)
}

func (this *MyStack) Pop() int {
    val := this.queue[0]
    this.queue = this.queue[1:]
    return val
}

func (this *MyStack) Top() int {
    return this.queue[0]
}

func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}
