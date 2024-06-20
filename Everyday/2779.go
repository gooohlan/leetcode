package Everyday

import "sort"

func maximumBeauty(nums []int, k int) int {
    n := len(nums)
    sort.Ints(nums)
    
    d := make([]int, nums[n-1]+k+2)
    for _, num := range nums {
        d[max(0, num-k)] += 1
        d[num+k+1] -= 1
    }
    
    sum := 0
    ans := 0
    for i := max(0, nums[0]-k); i < len(d); i++ {
        sum += d[i]
        ans = max(ans, sum)
    }
    return ans
}

func maximumBeauty2(nums []int, k int) int {
    n := len(nums)
    sort.Ints(nums)
    
    left, ans := 0, 1
    for right := 1; right < n; right++ {
        if nums[right]-nums[left] > k*2 {
            left++
        }
        ans = max(ans, right-left+1)
    }
    
    return ans
}
