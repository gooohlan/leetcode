package DataStructure

func missingNumber(nums []int) int {
    res := len(nums)
    for i, num := range nums {
        res ^= i ^ num
    }
    return res
}
