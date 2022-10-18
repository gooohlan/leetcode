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

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}
