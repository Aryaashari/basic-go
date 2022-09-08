package main

import "fmt"

func main() {

	var x int

	fmt.Scanln(&x)

	for x >= 10 {
		fmt.Printf("%d ", x%10)
		x /= 10
	}

	fmt.Printf("%d", x)

}
