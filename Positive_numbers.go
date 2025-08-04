package main

import "fmt"

func main() {
	var n, counter int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		if number > 0 {
			counter++
		}
	}

	fmt.Println(counter)
}
