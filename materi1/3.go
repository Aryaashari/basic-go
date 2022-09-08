package main

import "fmt"

func main() {

	var diskon, cashback, kartu, agreeCreatCard bool
	var totalBelanja int

	fmt.Scanln(&totalBelanja, &agreeCreatCard)

	diskon = totalBelanja > 100000
	kartu = agreeCreatCard && diskon
	cashback = kartu && totalBelanja > 200000

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("CashBack?", cashback)
}
