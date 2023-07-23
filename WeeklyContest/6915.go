package WeeklyContest

// 2,3,7,9,3
// 5,7,9,3 -> 12,9,3 -> 5,16,3 -> 21,3
// 2,10,9,3 -> 12,9,3
// 2,3,16,3 -> 5, 16, 3 -> 2,19,3
// 2,10,5,3,5
func maxArrayValue(nums []int) int64 {
    for i := len(nums) - 1; i > 0; {
        if nums[i] >= nums[i-1] {
            nums[i] += nums[i-1]
            nums = append(nums[:i-1], nums[i:]...)
        }
        i--
    }
    res := 0
    for _, num := range nums {
        if num > res {
            res = num
        }
    }
    return int64(res)
}
