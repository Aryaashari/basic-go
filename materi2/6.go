package main

import "fmt"

func main() {

	var n int
	var topi, ikatPinggang, alatTulis, tali, pisau, status bool

	fmt.Scanln(&n)

	status = true

	for i := 0; i < n; i++ {
		if status {
			fmt.Scanln(&topi, &ikatPinggang, &alatTulis, &tali, &pisau)
			status = topi && ikatPinggang && alatTulis && tali && pisau
		} else {
			break
		}
	}

	fmt.Println(status)
}
