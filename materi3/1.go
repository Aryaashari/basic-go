package main

import "fmt"

func main() {
	var penumpang, kapasitas int

	fmt.Scanln(&penumpang, &kapasitas)

	if penumpang <= kapasitas {
		fmt.Println("1 bus")
	} else {
		hasilBagi := penumpang / kapasitas
		if penumpang%kapasitas != 0 {
			hasilBagi += 1
		}
		fmt.Println(hasilBagi, "bus")
	}
}
