package main

import "fmt"

func main() {
	fmt.Printf("Total : %d", sum())
}

func sum() int {
	var limit_number int
	var total int
	fmt.Print("Sınır sayısı : ")
	fmt.Scan(&limit_number)

	for i := 1; i <= limit_number; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	return total
}
