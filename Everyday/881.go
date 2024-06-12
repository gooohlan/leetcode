package Everyday

import "sort"

func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    n := len(people)
    left, right := 0, n-1
    ans := 0
    for left <= right {
        if people[left]+people[right] <= limit {
            left++
        }
        right--
        ans++
    }
    return ans
}
