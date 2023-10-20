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
