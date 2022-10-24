package BinarySearch

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 递归
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	mid := (r-l)/2 + l
	num := nums[mid]
	if num == target {
		return mid
	}
	if target < num {
		return search2(nums[:mid], target)
	}
	if target > num {
		rNum := search2(nums[mid+1:], target)
		if rNum != -1 {
			// 递归向右时往后移了一位,所以要+1
			return mid + rNum + 1
		}
	}
	// 未找到则返回-1
	return -1
}
