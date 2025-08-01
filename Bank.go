package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	x = x * 100
	y = y * 100
	for i := 1; x < y; i++ {
		x += x * p / 100
		if x >= y {
			fmt.Println(i)
		}
	}
}
