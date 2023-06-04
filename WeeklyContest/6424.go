package WeeklyContest

func semiOrderedPermutation(nums []int) int {
    n := len(nums)
    l, r := 0, 0
    for i := 0; i < n; i++ {
        if nums[i] == 1 {
            l = i
        }
        if nums[i] == n {
            r = i
        }
    }
    if r > l {
        return l + n - r - 1
    } else {
        return l + n - r - 1 - 1
    }
}
