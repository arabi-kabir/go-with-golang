package main

import "fmt"

func main() {
	// slice with range
	nums := []int{6, 2, 3, 4, 5}

	sum := 0
	for i, v := range nums {
		println(i, v)
		sum += v
	}
	fmt.Println("Sum:", sum)

	// map with range
	ages := map[string]int{"John": 30, "Jane": 25, "Doe": 40}

	for k, v := range ages {
		println(k, v)
	}

	// string with range
	name := "John Doe"

	for i, v := range name {
		println(i, string(v), v)
	}
}
