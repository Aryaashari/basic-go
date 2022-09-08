package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	if n%3 == 0 {
		fmt.Println("Kelipatan 3")
	}

	if n%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}
