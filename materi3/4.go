package main

import "fmt"

func main() {
	var uts, uas, kuis, praktikum uint
	var nilaiAkhir float64

	fmt.Print("Nilai UTS: ")
	fmt.Scanln(&uts)

	fmt.Print("Nilai UAS: ")
	fmt.Scanln(&uas)

	fmt.Print("Nilai Kuis: ")
	fmt.Scanln(&kuis)

	fmt.Print("Nilai Praktikum: ")
	fmt.Scanln(&praktikum)

	if uts <= 100 && uas <= 100 && kuis <= 100 && praktikum <= 100 {
		nilaiAkhir = (35 * float64(uts) / 100) + (35 * float64(uas) / 100) + (20 * float64(kuis) / 100) + (10 * float64(praktikum) / 100)
		fmt.Println("Nilai Akhir:", nilaiAkhir)

		if nilaiAkhir > 85 {
			fmt.Println("Indeks mutu: A")
		} else if nilaiAkhir > 70 {
			fmt.Println("Indeks mutu: B")
		} else if nilaiAkhir > 50 {
			fmt.Println("Indeks mutu: C")
		} else if nilaiAkhir > 40 {
			fmt.Println("Indeks mutu: D")
		} else {
			fmt.Println("Indeks mutu: E")
		}
	} else {
		fmt.Println("NILAI TIDAK VALID")
	}
}
