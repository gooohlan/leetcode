package WeeklyContest

func maximumJumps(nums []int, target int) int {
    n := len(nums)
    var count int
    for i, j := 0, 1; j < n; j++ {
        if abs(nums[i], nums[j]) <= target {
            count++
            i = j
        }
    }
    if count == 0 {
        return -1
    }
    return count
}

// func abs(a, b int) int {
//     if a > b {
//         return a - b
//     }
//     return b - a
// }
