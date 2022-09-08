package main

import "fmt"

func main() {

	var t, v, kapasitas int

	fmt.Scanln(&t)
	fmt.Scanln(&v)
	kapasitas = v
	for kapasitas < t {
		fmt.Printf("%t \n", false)
		fmt.Scanln(&v)
		kapasitas += v
	}

	fmt.Printf("%t", true)

}
