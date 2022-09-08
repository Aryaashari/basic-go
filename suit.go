package main

import "fmt"

func main() {
	var n, nPutaran int
	var p1, p2 string

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&nPutaran)

		poinP1 := 0
		poinP2 := 0

		for j := 0; j < nPutaran; j++ {
			fmt.Scanln(&p1, &p2)
			if p1 != p2 {
				if p1 == "B" && p2 == "G" || p1 == "G" && p2 == "K" || p1 == "K" && p2 == "B" {
					poinP1 += 1
				} else {
					poinP2 += 1
				}
			}
		}

		if poinP1 > poinP2 {
			fmt.Println("Pemain 1")
		} else if poinP1 < poinP2 {
			fmt.Println("Pemain 2")
		} else {
			fmt.Println("SERI")
		}
	}
}
