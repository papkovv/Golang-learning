package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var slice []int

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		slice = append(slice, number)
	}

	fmt.Println(slice[3])
}
