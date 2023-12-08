package main

import "fmt"

func main() {
	fmt.Printf("Sonuç : %d", factorial(5))
}

func factorial(number int) int {
	var result = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}
