package main

import "fmt"

func main() {
	var n, n1, n2, n3, summ int

	fmt.Scan(&n)

	n1 = n / 100
	n2 = (n / 10) % 10
	n3 = n % 10
	summ = n1 + n2 + n3

	fmt.Println(summ)
}
