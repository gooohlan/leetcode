package Everyday

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if j-i >= indexDifference && abs(nums[i], nums[j]) >= valueDifference {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}
}

func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func findIndices2(nums []int, indexDifference int, valueDifference int) []int {
    minI, maxI := 0, 0
    for j := indexDifference; j < len(nums); j++ {
        i := j - indexDifference
        if nums[i] < nums[minI] {
            minI = i
        }
        if nums[i] > nums[maxI] {
            maxI = i
        }
        
        if nums[j]-nums[minI] >= valueDifference {
            return []int{minI, j}
        }
        if nums[maxI]-nums[j] >= valueDifference {
            return []int{maxI, j}
        }
    }
    return []int{-1, -1}
}
