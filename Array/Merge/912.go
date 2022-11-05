package Merge

// 归并排序, 取一个中间值,先排左边,再排右边,然后合并
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	l, r := 0, len(nums)-1
	mid := l + (r-l)/2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid+1:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	ln, rn := len(left), len(right)
	tmp := make([]int, ln+rn-1)

	var cnt int
	i, j := 0, 0
	for i < ln && j < ln {
		if left[i] <= right[j] {
			tmp[cnt] = left[i]
			i++
		} else {
			tmp[cnt] = right[j]
			j++
		}
		cnt++
	}

	if i < ln {
		tmp[cnt] = left[i]
		i++
		cnt++
	}

	if j < ln {
		tmp[cnt] = right[j]
		j++
		cnt++
	}
	return tmp
}
