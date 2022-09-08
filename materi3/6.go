package main

import "fmt"

func main() {
	const dewasa int = 3
	const kecil int = 5
	const kapasitasDewasa int = 5
	const kapasitasKecil int = 2

	var orang int

	fmt.Scanln(&orang)

	max := (dewasa * kapasitasDewasa) + (kecil * kapasitasKecil)

	if orang > max {
		fmt.Printf("Dewasa 3, Kecil 5, dan %d tidak berangkat", orang-max)
	} else if orang == max {
		fmt.Printf("Dewasa 3, Kecil 5")
	} else {
		if orang <= dewasa*kapasitasDewasa {
			if orang%kapasitasDewasa == 0 {
				fmt.Printf("Dewasa %d", orang/kapasitasDewasa)
			} else {
				fmt.Printf("Dewasa %d", orang/kapasitasDewasa+1)
			}
		} else {
			if (orang-(dewasa*kapasitasDewasa))%kapasitasKecil == 0 {
				fmt.Printf("Dewasa 3, Kecil %d", (orang-(dewasa*kapasitasDewasa))/kapasitasKecil)
			} else {
				fmt.Printf("Dewasa 3, Kecil %d", (orang-(dewasa*kapasitasDewasa))/kapasitasKecil+1)
			}
		}
	}

}
