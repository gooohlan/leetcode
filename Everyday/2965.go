package Everyday

func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    cnt := make([]int, n*n+1)
    for _, g := range grid {
        for _, x := range g {
            cnt[x]++
        }
    }
    
    ans := make([]int, 2)
    for i := 1; i <= n*n; i++ {
        if cnt[i] == 0 {
            ans[1] = i
        } else if cnt[i] == 2 {
            ans[0] = i
        }
    }
    
    return ans
}
