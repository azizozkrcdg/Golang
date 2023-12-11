package main

import "fmt"

func main() {
	fmt.Printf("Sonuç : %d", exponential_num_calculate(4, 3))
}

func exponential_num_calculate(number int, force int) int {
	var result = 1
	for i := 1; i <= force; i++ {
		result *= number
	}
	return result
}
