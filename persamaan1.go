package main

import "fmt"

func f(x float64) float64 {
	return (x * x) - (6 * x) + 9
}

func main() {
	var fCurrent, fLast, fNext float64

	fmt.Print("X0: ")
	fmt.Scanln(&fLast)

	fmt.Print("X1: ")
	fmt.Scanln(&fCurrent)

	for i := 2; i <= 11; i++ {
		fNext = fCurrent - (f(fCurrent) * (fLast - fCurrent) / (f(fLast) - f(fCurrent)))
		fmt.Printf("X%d: %g \n", i, fNext)
		fLast = fCurrent
		fCurrent = fNext
	}

}
