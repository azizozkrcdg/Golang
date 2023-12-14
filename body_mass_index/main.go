package main

import "fmt"

func main() {
	var body_mass_index = calculate()
	fmt.Printf("Vücut kitle indeksi : %f", body_mass_index)
	fmt.Println()

	if body_mass_index <= 18.5 {
		fmt.Println("İdeal kilonun altında.")
	} else if 18.5 <= body_mass_index && body_mass_index <= 24.9 {
		fmt.Println("İdeal kilo")
	} else if 25.0 <= body_mass_index && body_mass_index <= 29.9 {
		fmt.Println("İdeal kilonun üstünde")
	} else if 30.0 <= body_mass_index && body_mass_index <= 39.9 {
		fmt.Println("İdeal kilonun çok üstünde (Obez)")
	} else {
		fmt.Println("İdeal kilonun çok üstünde (Morbid Obez)")
	}
}

func calculate() float32 {
	var (
		size   float32
		kg     float32
		result float32
	)

	fmt.Print("Boy (örn. 1.75): ")
	fmt.Scan(&size)
	fmt.Print("Kilo : ")
	fmt.Scan(&kg)

	result = kg / (float32(size * size))

	return result
}
