package WeeklyContest

func countCompleteSubarrays(nums []int) int {
    set := make(map[int]struct{})
    for _, num := range nums {
        set[num] = struct{}{}
    }
    l := len(set)
    
    cnt := make(map[int]int)
    left := 0
    res := 0
    for _, num := range nums {
        cnt[num]++
        for len(cnt) == l { // 满足条件,左端点向右移动
            x := nums[left]
            cnt[x]--
            if cnt[x] == 0 {
                delete(cnt, x)
            }
            left++
        }
        res += left
    }
    return res
}
