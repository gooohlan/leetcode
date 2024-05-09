package WeeklyContest

func sumImbalanceNumbers(nums []int) int {
    n := len(nums)
    res := 0
    vis := make([]int, n+2)
    for i, _ := range vis {
        vis[i] = -1
    }
    for i, num := range nums {
        vis[num] = i
        cnt := 0
        for j := i + 1; j < n; j++ {
            x := nums[j]
            if vis[x] != i {
                cnt++
                if vis[x-1] == i {
                    cnt--
                }
                if vis[x+1] == i {
                    cnt--
                }
                vis[x] = i
            }
            res += cnt
        }
    }
    return res
}

func sumImbalanceNumbers2(nums []int) (res int) {
    n := len(nums)
    right := make([]int, n)
    idx := make([]int, n+1)
    for i := range idx {
        idx[i] = n
    }
    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        right[i] = min(idx[x], idx[x-1])
        idx[x] = i
    }
    
    for i := range idx {
        idx[i] = -1
    }
    for i, num := range nums {
        res += (i - idx[num-1]) * (right[i] - i)
        idx[num] = i
    }
    
    return res - n*(n+1)/2
}
