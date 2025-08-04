package main

import "fmt"

func main() {
	var n, min, counter int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		if i == 0 || number < min {
			min = number
			counter = 1
		} else if number == min {
			counter++
		}
	}

	fmt.Println(counter)
}
