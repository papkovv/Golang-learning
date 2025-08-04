package main

import "fmt"

func main() {
	var n, sum, result int

	fmt.Scan(&n)

	for n >= 10 {
		result = 0
		for n > 0 {
			sum = n % 10
			n /= 10
			result += sum
		}
		n = result
	}

	fmt.Println(result)
}
