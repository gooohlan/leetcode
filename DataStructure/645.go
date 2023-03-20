package DataStructure

func findErrorNums(nums []int) []int {
    n := len(nums)
    dup := -1
    for i := 0; i < n; i++ {
        index := abs(nums[i]) - 1
        if nums[index] < 0 {
            dup = abs(nums[i])
        } else {
            nums[index] *= -1
        }
    }
    
    missing := -1
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            missing = i + 1
        }
    }
    return []int{dup, missing}
}

func abs(i int) int {
    if i > 0 {
        return i
    }
    return -i
}
