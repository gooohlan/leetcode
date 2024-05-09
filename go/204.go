package main

func countPrimes(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes2(n int) int {
	var count int
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := i * 2; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return count
}

func countPrimes3(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, prime := range primes {
			if i*prime >= n {
				break
			}
			isPrime[i*prime] = false
			if i%prime == 0 {
				break
			}
		}
	}
	return len(primes)
}
