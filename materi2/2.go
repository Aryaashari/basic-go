package main

import "fmt"

func main() {
	var percobaan int
	var username, password string

	percobaan = 0

	fmt.Scanln(&username, &password)

	for username != "admin" && password != "admin" {
		percobaan += 1
		fmt.Scanln(&username, &password)
	}

	fmt.Printf("%d Login berhasil", percobaan)
}
