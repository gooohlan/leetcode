package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		num := nums[mid]
		if num == target {
			return mid
		} else if target < num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := (r-l)/2 + l
	num := nums[mid]
	if num == target {
		return mid
	} else if target < num {
		search2(nums[:r], target)
	} else {
		search2(nums[r:], target)
	}
	return -1
}
