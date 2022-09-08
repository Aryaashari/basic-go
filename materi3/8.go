package main

import "fmt"

func main() {
	var nominal int
	var countSeratus, countLimaratus, countSeribu, countDuaribu int
	countSeratus = 0
	countLimaratus = 0
	countSeribu = 0
	countDuaribu = 0

	fmt.Scanln(&nominal)

	for nominal == 100 || nominal == 500 || nominal == 1000 || nominal == 2000 {
		fmt.Scanln(&nominal)

		if nominal == 100 {
			countSeratus += 1
		} else if nominal == 500 {
			countLimaratus += 1
		} else if nominal == 1000 {
			countSeribu += 1
		} else if nominal == 2000 {
			countDuaribu += 1
		}
	}

	if countSeratus != 0 {
		fmt.Printf("Rp. 100 = %d keping \n", countSeratus)
	}
	if countLimaratus != 0 {
		fmt.Printf("Rp. 500 = %d keping \n", countLimaratus)
	}
	if countSeribu != 0 {
		fmt.Printf("Rp. 1000 = %d keping \n", countSeribu)
	}
	if countDuaribu != 0 {
		fmt.Printf("Rp. 2000 = %d keping \n", countDuaribu)
	}

}
