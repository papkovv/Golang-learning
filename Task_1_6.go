package main

import "fmt"

func main() {
	var a, b, c float32

	fmt.Scan(&a, &b)

	c = (a + b) / 2

	fmt.Println(c)
}
