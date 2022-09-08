package main

import "fmt"

func abc(x, y int) int {
	if x%y == 0 {
		return y
	} else {
		return abc(x, y-1)
	}
}

func bintang(x int) {
	if x != 0 {
		fmt.Println("*")
		x -= 1
		bintang(x)
		x += 1
	}
}

func cetak(n int) {
	for i := 1; i <= n+1; i++ {
		n += 1
	}

	if n%2 == 0 {
		fmt.Print("Ganjil")
	} else {
		fmt.Print("Genap")
	}
}

func dapdab(x int) {
	for i := 1; i <= (x % 9); i++ {
		fmt.Print("dab ")
		for j := 1; j <= (365 % (i * 2)); j++ {
			fmt.Print("dap ")
		}
	}
}

func f(n int) int {
	if n <= 1 {
		return 5
	}

	return n + f(n-1) + f(n-2)
}

func otakatik(j, k int) {
	for i := 1; i <= 10; i++ {
		if j < 0 {
			if k > 0 {
				j = j + k
			} else {
				k = (-j) + k
			}
		} else {
			if k < 0 {
				j = (-j) - k
			} else {
				k = j - k
			}
		}
	}

	fmt.Print(j, k)
}

func main() {
	// var x int

	// fmt.Println(abc(19, 9))
	// bintang(5)

	// fmt.Sscanf("Nilai X: ", "%d", &x)

	otakatik(5, 2)

	// for i := 1; i <= (365 % 8); i++ {
	// 	fmt.Println("dab")
	// 	for j := 1; j <= 365%(i*2); j++ {
	// 		fmt.Println("DAP")
	// 	}
	// }
}
