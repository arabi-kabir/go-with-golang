package main

import "fmt"

func main() {
	var name string = "John"

	var age int = 30

	var isStudent bool = true

	var height float64 = 5.9

	var weight float32 = 70.5

	var address string = "123 Main St"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Height:", height)
	fmt.Println("Weight:", weight)
	fmt.Println("Address:", address)

	// infer
	var name2 = "Alice"
	fmt.Println("Name2:", name2)

	// short
	name3 := "Bob"
	fmt.Println("Name3:", name3)

	// assign later
	var cgpa float64
	cgpa = 3.8
	fmt.Println("CGPA:", cgpa)
}
