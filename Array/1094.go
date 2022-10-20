package Array

// 依旧使用差分技巧, 求出每个车站人数,当车站人数超过capacity,表示超载
func carPooling(trips [][]int, capacity int) bool {
	end := 0
	diff := make([]int, 1001)
	for _, trip := range trips {
		if trip[2] > end {
			end = trip[2]
		}
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}
	if diff[0] > capacity {
		return false
	}
	for i := 1; i < end; i++ {
		diff[i] += diff[i-1]
		if diff[i] > capacity {
			return false
		}
	}
	return true
}
