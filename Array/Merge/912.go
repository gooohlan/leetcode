package Merge

// 归并排序, 取一个中间值,先排左边,再排右边,然后合并
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	ln, rn := len(left), len(right)
	res := make([]int, ln+rn)

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
	// if i < ln {
	// 	res[cnt] = left[i]
	// 	i++
	// 	cnt++
	// }
	copy(res[cnt+len(left)-i:], right[j:])
	// if j < ln {
	// 	res[cnt] = right[j]
	// 	j++
	// 	cnt++
	// }
	return res
}
