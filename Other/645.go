package Other

import (
    "sort"
)

func findErrorNums(nums []int) []int {
    res := make([]int, 2)
    sort.Ints(nums)
    var pre int
    for _, num := range nums {
        if pre == num {
            res[0] = num
        } else if num-pre > 1 {
            res[1] = pre + 1
        }
        pre = num
    }
    if nums[len(nums)-1] != len(nums) {
        res[1] = len(nums)
    }
    return res
}

func findErrorNums2(nums []int) []int {
    xor := 0
    for _, v := range nums {
        xor ^= v
    }
    n := len(nums)
    for i := 1; i <= n; i++ {
        xor ^= i
    }
    lowbit := xor & -xor
    num1, num2 := 0, 0
    for _, v := range nums {
        if v&lowbit == 0 {
            num1 ^= v
        } else {
            num2 ^= v
        }
    }
    for i := 1; i <= n; i++ {
        if i&lowbit == 0 {
            num1 ^= i
        } else {
            num2 ^= i
        }
    }
    for _, v := range nums {
        if v == num1 {
            return []int{num1, num2}
        }
    }
    return []int{num2, num1}
}

func findErrorNumsHash(nums []int) []int {
    res := make([]int, 2)
    m := make(map[int]int, 0)
    for _, num := range nums {
        m[num]++
    }
    for i := 1; i <= len(nums); i++ {
        if v := m[i]; v == 2 {
            res[0] = i
        } else if v == 0 {
            res[1] = i
        }
    }
    return res
}
