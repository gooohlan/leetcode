package Everyday

func distributeCandies(n int, limit int) int {
    ans := 0
    for x := 0; x <= limit; x++ {
        for y := 0; y <= limit; y++ {
            if x+y > n {
                break
            }
            if n-x-y <= limit {
                ans++
            }
        }
    }
    return ans
}

func cal(n int) int {
    if n < 2 {
        return 0
    }
    return n * (n - 1) / 2
}

func distributeCandies2(n, limit int) int {
    return cal(n+2) - 3*cal(n-limit-1) + 3*cal(n-2*limit) - cal(n-3*limit-1)
}
