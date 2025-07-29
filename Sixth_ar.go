package main

import "fmt"

func main() {
	var a, h, m int
	fmt.Scan(&a)
	h = a / 30
	m = 2 * (a % 30)
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}
