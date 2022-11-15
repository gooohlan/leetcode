package Merge

import (
	"math/rand"
	"time"
)

// 归并排序, 取一个中间值,先排左边,再排右边,然后合并
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	res := make([]int, len(left)+len(right))
	l, r := 0, 0
	for i := range nums {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
		}
	}

	return res
}

// 快排
func sortArrayFast(nums []int) []int {
	shuffle(nums) // 随机打乱,避免最坏情况超时
	var quick func(nums []int, left, right int) []int

	quick = func(nums []int, left, right int) []int {
		if left > right {
			return nil
		}

		// 左右指针和主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			// 寻找小于主元的右边元素
			for i < j && nums[j] >= pivot {
				j--
			}
			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}
			// 交换i,j下标元素,使其到达正确区域
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换元素,元素左右通过left区分
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)  // i左边进行排序
		quick(nums, i+1, right) // i右边排序
		return nums
	}

	return quick(nums, 0, len(nums)-1)
}

func shuffle(nums []int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(nums); i++ {
		r := i + rand.Intn(len(nums)-i)
		nums[i], nums[r] = nums[r], nums[i]
	}
}
