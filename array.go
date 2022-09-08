package main

import "fmt"

type arrInt [100]int
type A [100]arrInt

func main() {
	// var T arrInt
	// var S A
	// var N int
	// T[0] = 1
	// T[1] = 2
	// S[0] = T
	// fmt.Println(S)
	// fmt.Scanln(&N)
	// fmt.Println(T[:N])
	// isiArray(&T, N)
	// fmt.Println(T[:N])

	// reverse(T, N)
	// polindrome(T, N)

	// maksMinArray(T, N)

	// insertionSortASC(T, N)

	// insertionSortDESC(T, N)

	// sum(T, N)
	pola(1, 4)
	fmt.Println(f(4))
}

func pola(i, n int) {

	if i <= n {
		fmt.Println(i * n)
		pola(i+1, n)
	} else {
		polaMenurun(i-2, n)
	}

}

func f(x int) int {
	if x == 1 {
		return 2
	} else {
		return f(x)*f(x) + 1
	}
}
func polaMenurun(i, n int) {
	if i > 0 {
		fmt.Println(i * n)
		polaMenurun(i-1, n)
	}
}

func isiArray(T *arrInt, n int) {
	var i int
	i = 0
	for i < n {
		fmt.Scan(&T[i])
		i = i + 1
	}
}

func reverse(T arrInt, N int) {
	var reverseT arrInt
	var i, j int
	i = N - 1
	j = 0
	for i >= 0 {
		reverseT[j] = T[i]
		i = i - 1
		j = j + 1
	}
	fmt.Println(reverseT[:N])
}

func polindrome(T arrInt, N int) {
	var i, j int
	var isPolindrome bool
	i = 0
	j = N - 1
	isPolindrome = true
	for i < N/2 {
		if T[i] != T[j] {
			isPolindrome = false
			break
		}
		i = i + 1
		j = j - 1
	}
	fmt.Println(isPolindrome)
}

func maksMinArray(T arrInt, N int) {
	var maks, min, i int
	i = 1
	maks = 0
	min = 0

	for i < N {
		if T[maks] < T[i] {
			maks = i
		}

		if T[min] > T[i] {
			min = i
		}

		i = i + 1
	}
	fmt.Println("Nilai Maks:", T[maks])
	fmt.Println("Nilai Min:", T[min])
}

func insertionSortASC(T arrInt, N int) {
	var min, tmp int
	min = 0

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {

			if T[min] > T[j] {
				min = j
			}

		}
		tmp = T[i]
		T[i] = T[min]
		T[min] = tmp
		min = i + 1
	}

	fmt.Println(T[:N])
}

func insertionSortDESC(T arrInt, N int) {
	var max, tmp int
	max = 0

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if T[max] < T[j] {
				max = j
			}
		}
		tmp = T[i]
		T[i] = T[max]
		T[max] = tmp

		max = i + 1
	}

	fmt.Println(T[:N])
}

// 89 27 36 1 23 70 104 529 245 65

// var total int = 0

// func sum(T arrInt, N int) {

// 	if N > 0 {
// 		total = total + T[N-1]
// 		sum(T, N-1)
// 	}
// 	fmt.Println(total)
// }
