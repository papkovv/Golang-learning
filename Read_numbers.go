package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n <= 100 {
		if n >= 10 {
			fmt.Println(n)
		}
		fmt.Scan(&n)
	}
}
