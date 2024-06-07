package Everyday

func maxOperations(nums []int) int {
    s := nums[0] + nums[1]
    ans := 1
    for i := 3; i < len(nums) && nums[i-1]+nums[i] == s; i += 2 {
        ans++
    }
    return ans
}
