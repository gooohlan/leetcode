package WeeklyContest

// 数组里只有一个k，那么我们可以想到，以k为中心向左右两边扩展求子数组。
// 先找到k在nums里的位置idx；
// 从idx开始向左移动，记录路上大于k和小于k的数量，我这里采用正负性来计算，小于k就-1，大于k就+1，并用哈希表left记录下来，初始化left[0]=1；
// 再从idx开始向右移动，也记录路上大于k和小于k的数量，然后根据当前的值sum，看left中有多少数x能使得：sum+x==0或sum+x==1，因为如果为偶数，则大于k的数要比小于k的数多1才是满足条件的答案

func countSubarrays(nums []int, k int) int {
	idx := 0
	for nums[idx] != k {
		idx++
	}
	left := make(map[int]int)
	left[0] = 1 // 记录k
	sum := 0
	// 遍历左侧记录出正负数
	for i := idx - 1; i >= 0; i-- {
		if nums[i] < k {
			sum--
		} else {
			sum++
		}
		left[sum]++
	}

	cnt := left[0] + left[1] // 左侧符合条件的子数组,left[0]表示奇数长度,left[1]表示偶数长度
	sum = 0
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] < k {
			sum--
		} else {
			sum++
		}
		cnt += left[-sum]  // 奇数长度 sum+x==0
		cnt += left[1-sum] // 偶数长度 sum+x==1
	}

	return cnt
}
