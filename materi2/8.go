package main

import "fmt"

func main() {
	var lastData, currentData, totalPertambahan float32
	var n int = -1

	lastData = -1
	fmt.Scan(&currentData)

	for currentData != lastData {
		lastData = currentData
		fmt.Scan(&currentData)
		totalPertambahan += currentData - lastData
		n += 1
	}

	fmt.Println(totalPertambahan)
	fmt.Printf("%.3f", totalPertambahan/float32(n))
}
