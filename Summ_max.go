package main

import "fmt"

func main() {
	var n, max, summ int
	fmt.Scan(&n)
	for n != 0 {
		if n > max {
			max = n
			summ = 1
		} else if n == max {
			summ++
		}
		fmt.Scan(&n)
	}
	fmt.Println(summ)
}
