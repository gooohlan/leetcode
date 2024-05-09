package WeeklyContest

func findPrimePairs(n int) [][]int {
    var primes []int
    isP := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        isP[i] = true
    }
    for i := 2; i <= n; i++ {
        if isP[i] {
            primes = append(primes, i)
            for j := i * i; j <= n; j += i {
                isP[j] = false
            }
        }
    }
    
    //  如果n是奇数,因为只有奇数+偶数=奇数,而且偶数中只有2是质数,所以如果n是奇数是,最多只有一个质数对(2,n-2)
    if n%2 > 0 {
        if n > 4 && isP[n-2] {
            return [][]int{{2, n - 2}}
        }
    }
    var res [][]int
    for _, x := range primes {
        y := n - x
        if x <= y && isP[y] {
            res = append(res, []int{x, y})
        }
    }
    return res
}
