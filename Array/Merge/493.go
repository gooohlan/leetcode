package Merge

// 依旧是有归并排序的思路作为主框架
// 根据题意,我们再合并时判断一下题目条件即可
// 对于每个左边的元素,我们都去右边寻找符合条件的元素

var count int

func sort493(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := sort493(nums[:mid])
	right := sort493(nums[mid:])
	return merge493(left, right)
}

func reversePairs(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sort493(nums)
	return count
}

func merge493(left, right []int) []int {
	ln, rn := len(left), len(right)
	res := make([]int, ln+rn)

	// for i := 0; i < ln; i++ {
	// 	for j := 0; j < rn; j++ {
	// 		if left[i] > right[j]*2 {
	// 			count++
	// 		}
	// 	}
	// }
	// 合并时传递进来的数组都是排好序的,nums[i] <= nums[i+1]
	// 所以 当我们找到nums[i] > 2*nums[j]的nums[j],也必然也符合nums[i+1] > 2*nums[j]
	j := 0 // j用来记录右数组符合添加个数
	for i := 0; i < ln; i++ {
		for j < rn && left[i] > right[j]*2 { // 符合条件时右侧数组指针后移
			j++
		}
		count += j
	}

	var cnt int
	i, j := 0, 0
	for i < ln && j < ln {
		if left[i] <= right[j] {
			res[cnt] = left[i]
			i++
		} else {
			res[cnt] = right[j]
			j++
		}
		cnt++
	}

	copy(res[cnt:], left[i:]) // 复制左边为使用的
	copy(res[cnt+len(left)-i:], right[j:])
	return res
}
