package Other

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    res := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++ {
        j := 0
        for ; j < len(nums2) && nums1[i] != nums2[j]; {
            j++
        }
        for ; j < len(nums2); j++ {
            if nums1[i] < nums2[j] {
                res[i] = nums2[j]
                break
            }
        }
        if j == len(nums2) {
            res[i] = -1
        }
    }
    return res
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
    stack := make([]int, 0)
    mapList := make(map[int]int)
    for _, v := range nums2 {
        if len(stack) == 0 {
            stack = append(stack, v)
            continue
        }
        i := 0
        for ; i < len(stack); i++ {
            if v > stack[i] {
                break
            }
        }
        for j := i; j < len(stack); j++ {
            mapList[stack[j]] = v
        }
        stack = stack[:i]
        stack = append(stack, v)
    }
    for _, v := range stack {
        mapList[v] = -1
    }
    for k, v := range nums1 {
        nums1[k] = mapList[v]
    }
    return nums1
}
