package main

import "fmt"

func main() {
	var n int
	var bunga, pita string
	pita = ""

	fmt.Print("N: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga ke-%d: ", i)
		fmt.Scanln(&bunga)

		pita += bunga

		if i != n {
			pita += "-"
		}
	}

	fmt.Printf("Pita: %s", pita)
}
