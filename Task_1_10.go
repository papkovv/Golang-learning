package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	mod := ((b % 7) + 7) % 7
	max := b - mod

	if max >= a {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}
