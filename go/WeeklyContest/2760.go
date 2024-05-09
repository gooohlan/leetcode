package WeeklyContest

func longestAlternatingSubarray(nums []int, threshold int) int {
    right, n := 0, len(nums)
    var res int
    for right < n {
        if nums[right]%2 > 0 || nums[right] > threshold {
            right++
        } else {
            left := right
            right++
            for ; right < n && nums[right] <= threshold && nums[right]%2 != nums[right-1]%2; right++ {
            }
            res = max(res, right-left)
        }
    }
    return res
}
