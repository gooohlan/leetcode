package Array

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	return NumArray{
		sums: sums,
	}
}

func (nums *NumArray) SumRange(left int, right int) int {
	return nums.sums[right+1] - nums.sums[left]
}
