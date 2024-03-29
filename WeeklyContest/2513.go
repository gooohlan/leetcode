package WeeklyContest

import "sort"

// 假设第n个不能被d整除的数是m
// 在[0,m]中共有m/d个数被d整除了m,又因为0-m中有m+1个数,因此可知 n= m+1 - m/d
// m = (n-1) * d / (d-1),因为上述例子从0开始,要从正整数算,所以 m =  (n-1) * d / (d-1)+1
func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
    lcm := divisor1 * divisor2 / gcd(divisor1, divisor2)
    res1 := (uniqueCnt1-1)*divisor1/(divisor1-1) + 1
    res2 := (uniqueCnt2-1)*divisor2/(divisor2-1) + 1
    res3 := lcm*(uniqueCnt1+uniqueCnt2-1)/(lcm-1) + 1
    return max(max(res1, res2), res3)
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}

func minimizeSet2(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
    res1 := uniqueCnt1 + (uniqueCnt1-1)/(divisor1-1)
    res2 := uniqueCnt2 + (uniqueCnt2-1)/(divisor2-1)
    lcm := divisor1 * divisor2 / gcd(divisor1, divisor2)
    res3 := uniqueCnt1 + uniqueCnt2 + (uniqueCnt1+uniqueCnt2)/(lcm-1)
    return max(max(res1, res2), res3)
}

// 二分
func minimizeSet3(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
    lcm := divisor1 * divisor2 / gcd(divisor1, divisor2)
    
    // 最坏情况, d1=d2=2,全部取奇数
    return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
        left1 := max(uniqueCnt1-x/divisor2+x/lcm, 0) // 能被d2整除,且不能被最小公约数整除
        left2 := max(uniqueCnt2-x/divisor1+x/lcm, 0) // 能被d1整除,且不能被最小公约数整除
        common := x - x/divisor1 - x/divisor2 + x/lcm
        return common >= left1+left2
    })
}
