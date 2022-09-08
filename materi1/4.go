package main

import "fmt"

func main() {
	var digit, d1, d2, d3 int

	fmt.Scanln(&digit)

	d3 = digit % 10
	digit /= 10

	d2 = digit % 10
	digit /= 10

	d1 = digit % 10

	fmt.Println(d1, d2, d3)

}
