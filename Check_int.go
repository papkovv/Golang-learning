package main

import "fmt"

func main() {
	var i, a, b, c int
	fmt.Scan(&i)
	a = i / 100
	b = (i / 10) % 10
	c = i % 10
	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
