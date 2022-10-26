package BinarySearch

func splitArray(nums []int, k int) int {
	l, r := 0, 0
	for _, num := range nums {
		if num > l {
			l = num
		}
		r += num
	}

	for l < r {
		mid := int(uint(l+r) >> 1)
		if f401(nums, mid) <= k { // 缩小max 让数组数量增多达到k
			r = mid
		} else { // 加大
			l = mid + 1
		}
	}
	return l
}

// 最大子数组和为max时,nums 至少可以被分割成几个子数组
func f401(nums []int, max int) int {
	count := 1 // 至少分配数组个数
	sum := 0   // 子数组元素和
	for _, num := range nums {
		if sum+num > max {
			// 子数组和超过max 无法继续添加
			count++
			sum = num
		} else {
			// 可以继续添加
			sum += num
		}
	}

	return count
}
