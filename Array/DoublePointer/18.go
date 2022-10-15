package Array

import "sort"

// 在三数之和上再套一层循环
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		item := threeSum18(nums[i+1:], target-nums[i], nums[i])
		res = append(res, item...)
	}
	return res
}

func threeSum18(nums []int, target, first int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2 && nums[i]+nums[i+1]+nums[i+2] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		item := twoSum18(nums[i+1:], target-nums[i], first, nums[i])
		res = append(res, item...)
	}
	return res
}

func twoSum18(numbers []int, target, first, second int) [][]int {
	var res [][]int
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] < target {
			i++
			continue
		} else if numbers[i]+numbers[j] > target {
			j--
			continue
		} else {
			res = append(res, []int{first, second, numbers[i], numbers[j]})
			i++
			j--
		}

		// 去除头尾相同的数字
		for i < j && numbers[i] == numbers[i-1] {
			i++
		}
		for i < j && numbers[j] == numbers[j+1] {
			j--
		}
	}
	return res
}
