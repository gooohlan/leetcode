package WeeklyContest

// 统计满足x= nums1[i] = nums2[i]的数对个数swapCnt,以及x的众数mode及起出现次数modeCnt
// 分类讨论
// 1 如果modeCnt没有超过swapCnt的一半:
//   如果swapCnt是偶数,两两交换即可
//   如果swapCnt是奇数,那么只是有三种不同的x,其中一个必然可以和nums1[0]交换
//   这种情况下,代价就是这些x的下标之和
// 2 如果modeCnt超过swapCnt的一半(或者说modeCnt*2大于swapCnt):
//   根据鸠巢原理,无法通过重排这些数字,让数组不相等(因为还存在一些mode仍然相同)
//   此时必须不断寻找其他满足nums1[j]!=nums2[j]的数对,且数对中的数都不等于mode,直到满足条件1
// 为了让答案劲量小,应从左到右遍历数组,如果仍然无法满足,返回-1
func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	var ans int64
	var swapCnt, modeCnt, mode int   // 相同数对个数, 众数出现次数, 众数
	cnt := make([]int, len(nums1)+1) // 记录
	for i, x := range nums1 {
		if x == nums2[i] {
			ans += int64(i)
			swapCnt++
			cnt[x]++
			if cnt[x] > modeCnt {
				modeCnt, mode = cnt[x], x
			}
		}
	}

	for i, x := range nums1 {
		if modeCnt*2 <= swapCnt { // 众数出现次数没有超过数对的一半
			break
		}
		if x != nums2[i] && x != mode && nums2[i] != mode { // 满足nums1[j]!=nums2[j]的数对,且数对中的数都不等于mode
			ans += int64(i) // 交换时添加下标,并更新swapCnt的值,直到满足条件1
			swapCnt++
		}
	}
	if modeCnt*2 > swapCnt {
		return -1
	}
	return ans
}
