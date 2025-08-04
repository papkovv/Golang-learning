package main

import "fmt"

func main() {
	var k, h, m int

	fmt.Scan(&k)

	h = k / 3600
	m = (k - (h * 3600)) / 60

	fmt.Println("It is", h, "hours", m, "minutes.")
}
