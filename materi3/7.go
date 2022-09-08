package main

import "fmt"

func main() {
	var n int
	var isPrime bool = true

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if (i != 1 && i != n) && n%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Println(isPrime)
}
