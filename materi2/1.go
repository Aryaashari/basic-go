package main

import "fmt"

func main() {

	var u, n int

	fmt.Scanln(&u, &n)

	total := 0

	for i := 0; i < n; i++ {
		total += u
	}

	fmt.Println(total)
}
