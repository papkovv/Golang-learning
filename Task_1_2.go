package main

import "fmt"

func main() {
	var n, n1, n2, n3, new_n int

	fmt.Scan(&n)

	n1 = n / 100
	n2 = (n / 10) % 10
	n3 = n % 10
	new_n = n3*100 + n2*10 + n1

	fmt.Println(new_n)
}
