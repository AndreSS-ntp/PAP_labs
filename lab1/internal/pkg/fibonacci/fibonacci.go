package fibonacci

import "fmt"

func GiveFibonacci(n int) {
	a, b := 0, 1

	fmt.Printf("0 ")
	for b < n {
		fmt.Printf("%d ", b)
		a, b = b, a+b
	}
	fmt.Printf("\nFibonacci to num %d\n", n)
}
