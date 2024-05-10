package Other

import "math"

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

type MinStack struct {
    stack    []int
    minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack:    []int{},
        minStack: []int{math.MaxInt64},
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    this.minStack = append(this.minStack, min(val, this.minStack[len(this.minStack)-1]))
}

func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
    this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
