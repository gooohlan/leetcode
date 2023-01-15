package WeeklyContest

// 双指针,向右移动找到大于 k 对重复数组,然后移动左指针缩小窗口
// 由于左端点及其左边的都可以是好子数组的左端点，所以每个右端点对应的答案个数为 left+1
func countGood(nums []int, k int) int64 {
    cnt := make(map[int]int)
    left, pairs, res := 0, 0, 0
    for _, num := range nums {
        pairs += cnt[num]
        cnt[num]++
        
        for left < len(nums) && pairs-cnt[nums[left]]+1 >= k {
            pairs -= cnt[nums[left]] - 1
            cnt[nums[left]]--
            left++
        }
        
        if pairs >= k {
            res += left + 1
        }
    }
    return int64(res)
}
