package main

import "fmt"

func main() {
	var a, n1, n2, summ, counter int
	fmt.Scan(&a)

	n1 = 1
	n2 = 1
	counter = 2

	for summ < a {
		summ = n1 + n2
		n1 = n2
		n2 = summ
		counter++
	}

	if summ == a {
		fmt.Println(counter)
	} else {
		fmt.Println(-1)
	}
}
