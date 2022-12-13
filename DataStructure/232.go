package DataStructure

type MyQueue struct {
    stack []int
}

func Constructor232() MyQueue {
    return MyQueue{
        stack: make([]int, 0),
    }
}

func (this *MyQueue) Push(x int) {
    this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
    v := this.stack[0]
    this.stack = this.stack[1:]
    return v
}

func (this *MyQueue) Peek() int {
    return this.stack[0]
}

func (this *MyQueue) Empty() bool {
    return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
