package DataStructure

func countPrimes(n int) int {
    isPrime := make([]bool, n)
    for i := 2; i*i < n; i++ {
        if !isPrime[i] {
            for j := i * i; j < n; j += i {
                isPrime[j] = true
            }
        }
    }
    
    res := 0
    for i := 2; i < n; i++ {
        if !isPrime[i] {
            res++
        }
    }
    return res
}
