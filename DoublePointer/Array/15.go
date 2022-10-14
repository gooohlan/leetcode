package Array

import "sort"

// 假设结果第一个数为k, 那么另外两个数就为 0-k,这不就又变成两数之和吗
// 注意去掉头尾的相同数字
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		item := twoSum15(nums[i+1:], 0-nums[i])
		res = append(res, item...)
	}
	return res
}

func twoSum15(numbers []int, target int) [][]int {
	var res [][]int
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] < target {
			i++
			continue
		} else if numbers[i]+numbers[j] > target {
			j--
			continue
		} else {
			res = append(res, []int{target * -1, numbers[i], numbers[j]})
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
