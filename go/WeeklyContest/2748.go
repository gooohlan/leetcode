package WeeklyContest

func countBeautifulPairs(nums []int) int {
    var res int
    for i, x := range nums {
        for x >= 10 {
            x /= 10
        }
        for j := i + 1; j < len(nums); j++ {
            y := nums[j] % 10
            if gcd(x, y) == 1 {
                res++
            }
        }
    }
    return res
}

func countBeautifulPairs1(nums []int) int {
    cnt := [10]int{}
    res := 0
    for _, x := range nums {
        for y := 1; y < 10; y++ {
            if cnt[y] > 0 && gcd(x%10, y) == 1 { // 这里y 表示最高位,出现过最高位,且与最低位互质
                res += cnt[y]
            }
        }
        for x >= 10 {
            x /= 10
        }
        cnt[x]++
    }
    return res
}
