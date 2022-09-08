package main

import "fmt"

func main() {
	var isFlush, isStraight bool
	var k1, k2, k3, k4, k5 string

	fmt.Scanln(&k1, &k2, &k3, &k4, &k5)

	isFlush = string(k1[len(k1)-1]) == string(k2[len(k2)-1]) && string(k2[len(k2)-1]) == string(k3[len(k3)-1]) && string(k3[len(k3)-1]) == string(k4[len(k4)-1]) && string(k4[len(k4)-1]) == string(k5[len(k5)-1])

	isStraight = int(k1[0])+1 == int(k2[0]) && int(k2[0])+1 == int(k3[0]) && int(k3[0])+1 == int(k4[0]) && int(k4[0])+1 == int(k5[0])

	fmt.Println("Straight?", isStraight)
	fmt.Println("Flush?", isFlush)
}
