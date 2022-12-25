package WeeklyContest

import "sort"

func maximumTastiness(price []int, k int) int {
    sort.Ints(price)
    var check func(int) bool
    
    check = func(x int) bool {
        cnt := 1 // 记录取到的数字
        prev := price[0]
        for _, p := range price[1:] {
            if p-prev >= x { // 两次取的差应该大于等于x
                cnt++
                prev = p
            }
        }
        return cnt >= k
    }
    l, r := 0, price[len(price)-1]
    for l+1 < r {
        mid := (r-l)/2 + l
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }
    return l
}
