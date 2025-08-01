package main

import "fmt"

func main() {

	var a, b, counter_a, counter_b, number_a, number_b, step_a, step_b, final, final_c int

	fmt.Scan(&a)
	fmt.Scan(&b)

	switch {
	case a < 10:
		counter_a = 1
	case a < 100:
		counter_a = 2
	case a < 1000:
		counter_a = 3
	case a < 10000:
		counter_a = 4
	case a == 10000:
		counter_a = 5
	default:
		counter_a = 1
	}

	switch {
	case b < 10:
		counter_b = 1
	case b < 100:
		counter_b = 2
	case b < 1000:
		counter_b = 3
	case b < 10000:
		counter_b = 4
	case b == 10000:
		counter_b = 5
	default:
		counter_b = 1
	}

	for i := 1; i <= counter_a; i++ {

		step_a = 1

		for j := 1; j < i; j++ {
			step_a *= 10
		}

		if i == 1 {
			number_a = a % 10
		} else {
			number_a = (a / step_a) % 10
		}

		for k := 1; k <= counter_b; k++ {

			step_b = 1

			for j := 1; j < k; j++ {
				step_b *= 10
			}

			if k == 1 {
				number_b = b % 10
			} else {
				number_b = (b / step_b) % 10
			}

			if number_a == number_b {
				if final_c == 0 {
					final = number_a
				} else {
					final = final*10 + number_a
				}
				final_c++
			}
		}
	}

	for i := 1; i <= final_c; i++ {

		step_a = 1

		for j := 1; j < i; j++ {
			step_a *= 10
		}

		if i == 1 {
			number_a = final % 10
		} else {
			number_a = (final / step_a) % 10
		}

		if i == 1 {
			fmt.Print(number_a)
		} else {
			fmt.Print(" ", number_a)
		}
	}
}
