package main

import "fmt"

func main() {
	var s1, s2, s3 int

	fmt.Scanln(&s1, &s2, &s3)

	if s1 == s2 && s2 == s3 {
		fmt.Println("Segitiga Sama Sisi")
	} else if s1 != s2 && s1 != s3 && s3 != s2 {
		fmt.Println("Segitiga Sembarang")
	} else {
		fmt.Println("Segitiga Sama Kaki")
	}
}
