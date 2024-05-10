package Array

func containsNearbyDuplicate(nums []int, k int) bool {
	mapArr := make(map[int]int)
	for key, val := range nums {
		if v, ok := mapArr[val]; ok {
			if key-v <= k {
				return true
			}
		}
		mapArr[val] = key
	}
	return false
}

// 维护一个长度为K+1的滑动窗口,窗口中存在重复元素则返回true
func containsNearbyDuplicate2(nums []int, k int) bool {
	mapArr := make(map[int]struct{})

	for i, num := range nums {
		if i > k {
			delete(mapArr, nums[i-k-1]) // 移除最左侧元素
		}

		if _, ok := mapArr[num]; ok {
			return true
		}

		mapArr[num] = struct{}{}
	}
	return false
}
