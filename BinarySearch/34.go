package BinarySearch

func searchRange(nums []int, target int) []int {
    l := leftSearch(nums, target)

    if l == len(nums) || nums[l] != target {
        return []int{-1, -1}
    }
    r := rightSearch(nums, target)
    return []int{l, r}
}

// 数组中不存在target时分两种情况
// target 比数组中所有元素都要大,此时l = len(nums)
// 数组查找到最后一个元素,为找到target, 此时nums[l] != target
func leftSearch(nums []int, target int) int {
    l, r := 0, len(nums)

    for l < r {
        mid := l + (r-l)/2
        if nums[mid] >= target { // nums[mid] == target 与 nums[mid] > target合并
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

func rightSearch(nums []int, target int) int {
    l, r := 0, len(nums)

    for l < r {
        mid := l + (r-l)/2
        if nums[mid] <= target { // nums[mid] == target 与 nums[mid] < target合并
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l - 1
}
