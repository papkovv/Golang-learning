package main

import "fmt"

func main() {
	var i, n1, n2, n3, n4, n5, n6 int
	fmt.Scan(&i)
	n1 = i / 100000
	n2 = (i / 10000) % 10
	n3 = (i / 1000) % 10
	n4 = (i / 100) % 10
	n5 = (i / 10) % 10
	n6 = i % 10
	if n1+n2+n3 == n4+n5+n6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
