package Everyday

func distributeCandies1103(candies int, n int) []int {
    ans := make([]int, n)
    for i := 0; candies > 0; i++ {
        ans[i%n] += min(candies, i+1)
        candies -= i + 1
    }
    return ans
}
