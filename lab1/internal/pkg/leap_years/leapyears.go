package leap_years

import "fmt"

func GiveLeapYears(from int, to int) {
	for i := from; i <= to; i++ {
		if i%400 == 0 {
			fmt.Printf("%d ", i)
		} else if i%4 == 0 && i%100 != 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\nLeap Years from %d to %d\n", from, to)
}
