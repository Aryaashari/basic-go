package main

import "fmt"

type club struct {
	id      string
	m, k, d int
}

const NMAX = 123

type arrClub [NMAX]club

func main() {
	var A arrClub
	var n int
	inputData(&A, &n)
	sorting(&A, &n)
	print(A, n)
}

func inputData(A *arrClub, n *int) {
	var c club
	var id string
	*n = 0

	fmt.Scan(&id)
	for id != "None" {
		c.id = id
		fmt.Scanln(&c.m, &c.k, &c.d)
		A[*n] = c
		*n++

		fmt.Scan(&id)
	}

}

func sorting(A *arrClub, n *int) {
	var totalMax, totalCurr int
	var tmp club

	max := 0

	for i := 0; i < *n; i++ {
		for j := i; j < *n; j++ {
			totalMax = (A[max].m * 4) + A[max].d
			totalCurr = (A[j].m * 4) + A[j].d
			if totalMax <= totalCurr {
				max = j
			}
		}
		tmp = A[i]
		A[i] = A[max]
		A[max] = tmp
		max = i + 1
	}
}

func print(A arrClub, n int) {
	if n >= 10 {
		for i := 0; i < 10; i++ {
			poin := (A[i].m * 4) + A[i].d
			fmt.Println(A[i].id, poin)
		}
	} else {
		for i := 0; i < n; i++ {
			poin := (A[i].m * 4) + A[i].d
			fmt.Println(A[i].id, poin)
		}
	}

}
