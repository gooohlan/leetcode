package WeeklyContest

func countCompleteSubarrays(nums []int) (res int) {
    set := make(map[int]struct{})
    for _, num := range nums {
        set[num] = struct{}{}
    }
    l := len(set)
    
    for i := 0; i < len(nums); i++ {
        cnt := make(map[int]int)
        for _, x := range nums[i:] {
            cnt[x]++
            if len(cnt) == l {
                res++
            }
        }
    }
    return res
}

func countCompleteSubarrays2(nums []int) (res int) {
    set := make(map[int]struct{})
    for _, num := range nums {
        set[num] = struct{}{}
    }
    l := len(set)
    n := len(nums)
    
    cnt := make(map[int]int)
    left := 0
    for right, num := range nums {
        cnt[num]++
        for len(cnt) == l { // 满足条件,左端点向右移动
            res += n - right // r右边有n-r-1个元素,加上自身就是n-r个元素
            x := nums[left]
            cnt[x]--
            if cnt[x] == 0 {
                delete(cnt, x)
            }
            left++
        }
    }
    return res
}
