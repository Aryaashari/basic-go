package main

import "fmt"

type student struct {
	name   string
	age    int
	gender string
	next   *student
}

func main() {
	student2 := student{
		name:   "Dul",
		age:    18,
		gender: "L",
		next:   nil,
	}

	student1 := student{
		name:   "Arya",
		age:    18,
		gender: "L",
		next:   &student2,
	}

	current := student1

	for current == student {
		fmt.Println(current.name)
		fmt.Println(current.age)
		fmt.Println(current.gender)
		current = *current.next
		fmt.Println(current)
	}

}
