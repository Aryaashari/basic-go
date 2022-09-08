package main

import (
	"fmt"
)

func main() {
	var n int
	var status bool = true
	fmt.Scanln(&n)

	for n >= 10 {
		digit := n % 10
		n = n / 10
		if digit-(n%10) != 1 && digit-(n%10) != -1 {
			status = false
			break
		}
	}

	fmt.Println(status)
}
