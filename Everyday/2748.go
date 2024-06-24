package Everyday

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func countBeautifulPairs(nums []int) int {
    ans := 0
    n := len(nums)
    
    for i := 0; i < n; i++ {
        for nums[i] >= 10 {
            nums[i] /= 10
        }
        for j := i + 1; j < n; j++ {
            if gcd(nums[i], nums[j]%10) == 1 {
                ans++
            }
        }
    }
    return ans
}

func countBeautifulPairs2(nums []int) int {
    ans := 0
    cnt := [10]int{}
    
    for _, num := range nums {
        for i := 1; i < 10; i++ {
            if cnt[i] > 0 && gcd(i, num%10) == 1 {
                ans += cnt[i]
            }
        }
        for num >= 10 {
            num /= 10
        }
        cnt[num]++ // 统计最高位的出现次数
    }
    return ans
}
