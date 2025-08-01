package main

import "fmt"

func main() {
	var n, number, summ int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number > 9 && number < 100 && number%8 == 0 {
			summ += number
		}
	}
	fmt.Println(summ)
}
