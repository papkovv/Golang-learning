package main

import "fmt"

func main() {
	var n, d, counter, sum, number, ten int

	fmt.Scan(&n, &d)

	for n > 0 {
		sum = n % 10
		n /= 10
		if sum != d {
			ten = 1

			for i := 0; i < counter; i++ {
				ten *= 10
			}

			number += sum * ten
		}
	}

	fmt.Println(number)
}
