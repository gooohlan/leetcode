package Array

type Solution398 struct {
    nums  []int
    index int
}

func Constructor398(nums []int) Solution398 {
    return Solution398{
        nums: nums,
    }
}

func (this *Solution398) Pick(target int) int {
    for {
        if this.index == len(this.nums) {
            this.index = 0
        }
        if this.nums[this.index] == target {
            this.index++
            return this.index - 1
        }
        this.index++
    }
}
