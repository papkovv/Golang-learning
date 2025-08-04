package main

import "fmt"

func main() {
	workArray := [10]uint8{}

	for idx := range workArray {
		fmt.Scan(&workArray[idx])
	}

	for i := 0; i < 3; i++ {
		var f_inx, s_inx, step uint8

		fmt.Scan(&f_inx)
		fmt.Scan(&s_inx)

		step = workArray[f_inx]
		workArray[f_inx] = workArray[s_inx]
		workArray[s_inx] = step
	}

	for idx := range workArray {
		fmt.Print(workArray[idx])
		fmt.Print(" ")
	}
}
