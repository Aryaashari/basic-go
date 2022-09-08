package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var bb, total float64
	total = 0
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&bb)
		total += bb
	}

	// fmt.Printf("%f", total/float64(n))
	fmt.Println(strconv.FormatFloat(total/float64(n), 'f', -1, 64), "KG")
}
