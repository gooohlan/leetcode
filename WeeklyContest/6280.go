package WeeklyContest

func closestPrimes(left int, right int) []int {
    var primes []int
    isPrim := make([]bool, right+1)
    for i := 2; i <= right; i++ {
        if !isPrim[i] {
            if i >= left {
                primes = append(primes, i)
            }
            for j := i * i; j <= right; j += i {
                isPrim[j] = true
            }
        }
    }
    
    p, q := -1, -1
    for i := 1; i < len(primes); i++ {
        if p < 0 || primes[i]-primes[i-1] < q-p {
            p, q = primes[i-1], primes[i]
        }
    }
    
    return []int{p, q}
}
