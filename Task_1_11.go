package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	lastDigit := n % 10
	lastTwo := n % 100

	var word string

	switch {
	case lastTwo >= 11 && lastTwo <= 14:
		word = "korov"
	case lastDigit == 1:
		word = "korova"
	case lastDigit >= 2 && lastDigit <= 4:
		word = "korovy"
	default:
		word = "korov"
	}

	fmt.Printf("%d %s", n, word)
}
