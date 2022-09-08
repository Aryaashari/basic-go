package main

import "fmt"

func main() {
	// var isKabisat bool
	var year int

	fmt.Scanln(&year)
	fmt.Println(year%400 == 0 || (year%4 == 0 && year%100 != 0))

}
