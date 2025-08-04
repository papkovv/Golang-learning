package main

import "fmt"

func main() {
	var n int
	var mas []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		mas = append(mas, number)
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(mas[i], " ")
		}
	}
}
