package main

import "fmt"

func main() {
	var n, two int
	fmt.Scan(&n)

	two = 1

	for two <= n {
		fmt.Print(two, " ")
		two = two * 2
	}
}
