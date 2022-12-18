package WeeklyContest

import (
    "math"
)

var primeNumbers = make(map[int]int)

func isPrime(n int) bool {
    
    if n > 2 && n%2 == 0 {
        primeNumbers[n] = 2
        return false
    }
    
    for i := 2; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
        if n%i == 0 {
            primeNumbers[n] = 2
            return false
        }
    }
    primeNumbers[n] = 1
    return true
}

func smallestValue(n int) int {
    primeNumbers = make(map[int]int)
    min := n
    for !isPrime(n) {
        newN := printPrime(n)
        if newN == n {
            return n
        }
        if min > newN {
            min = newN
        }
    }
    
    return min
}

func printPrime(n int) int {
    var sum int
    for j := 2; j < n; j++ {
        if isPrime(n) {
            return n
        }
        if n%j == 0 && isPrime(j) {
            n /= j
            sum += j
            j--
        }
    }
    return sum
}

func factor(n int) int {
    var sum int
    for i := n / 2; i >= 2 && n > 1; i-- {
        if isPrime(i) {
            if n%i == 0 {
                sum += i
                n = n / i
                i++
            }
        }
    }
    return sum
}
