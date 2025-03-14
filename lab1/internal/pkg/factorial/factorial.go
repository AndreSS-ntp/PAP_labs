package factorial

import "fmt"

func Factorial(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		num *= i
	}
	fmt.Printf("Factorial of %d: %d\n", n, num)
}
