package prime_nums

import "fmt"

func SieveOfEratosthenes(n int) {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var result []int
	for i := 2; i <= n; i++ {
		if primes[i] {
			result = append(result, i)
		}
	}
	fmt.Printf("Prime nums to %d: %d", n, result)
}
