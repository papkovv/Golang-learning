package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if i >= 0 && i < 10 {
		fmt.Println(i)
	} else if i >= 10 && i < 100 {
		fmt.Println(i / 10)
	} else if i >= 100 && i < 1000 {
		fmt.Println(i / 100)
	} else if i >= 1000 && i < 10000 {
		fmt.Println(i / 1000)
	} else if i == 10000 {
		fmt.Println(i / 10000)
	}
}
