package main

import "fmt"

func kuadrat(n int) int {
	return n * n
}

func main() {
	var k int
	var hasil float64

	fmt.Print("Nilai K: ")
	fmt.Scanln(&k)

	hasil = float64(kuadrat(4*k+2)) / (float64(4*k+1)*float64(4*k+3))

	fmt.Printf("Hasiil f(%d): %.10f", k, hasil)
}