package Merge

// 找出第K个最大的元素,就相当于升序排序后排名第n-k的元素,令k' = n - k
// 产考912的快排,交换元素时我们将中间值进行与k'进行比较,再去对应区域寻找
// 假设我们从大到小排序,那么第K个最大元素就是下标为k-1的元素
// 快排分割时可以当得知分割出下标,因此我们在更靠近K-1的区域继续寻找
func findKthLargest(nums []int, k int) int {
	k--
	l, r := 0, len(nums)-1
	for l <= r {
		p := partition(nums, l, r)
		if p == k {
			return nums[p]
		}
		if p > k {
			r = p - 1
		} else {
			l = p + 1
		}
	}
	return -1
}

func partition(nums []int, left, right int) int {
	// 左右指针和主元
	i, j, pivot := left, right, nums[right]
	for i < j {

		// 寻找左边小于于主元的元素
		for i < j && nums[i] >= pivot {
			i++
		}
		// 寻找右边大于主元的元素
		for i < j && nums[j] <= pivot {
			j--
		}

		// 交换i,j下标元素,使其到达正确区域
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
