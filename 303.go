package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	for i, v := range nums {
		sums[i] = v
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	var count int
	for i := left; i <= right; i++ {
		count += this.sums[i]
	}
	return count
}

func Constructor2(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num // sum[i+1]存储nums[i]的前缀和
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange2(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
