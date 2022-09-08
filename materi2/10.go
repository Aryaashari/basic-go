package main

import "fmt"

func main() {

	var n, f1, f2 int
	f1 = 0
	f2 = 1

	fmt.Scanln(&n)

	fmt.Print(f1, " ", f2)

	for i := 0; i < n-1; i++ {
		fn := f2 + f1
		fmt.Print(" ", fn)
		f1 = f2
		f2 = fn
	}

}
