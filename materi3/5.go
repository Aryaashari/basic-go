package main

import "fmt"

func main() {
	var n int = 1
	var currentDay, yesterday, penurunan int
	penurunan = 0

	fmt.Printf("Hari ke-%d: ", n)
	fmt.Scanln(&currentDay)

	for currentDay >= 0 {
		yesterday = currentDay
		n++

		fmt.Printf("Hari ke-%d: ", n)
		fmt.Scanln(&currentDay)

		if currentDay >= 0 {
			if currentDay > yesterday {
				fmt.Println("NAIK")
			} else if currentDay < yesterday {
				fmt.Println("TURUN")
				penurunan += 1
			} else {
				fmt.Println("STABIL")
			}
		}

	}

	fmt.Println("SELESAI")
	fmt.Printf("Terjadi penurunan %d kali", penurunan)
}
