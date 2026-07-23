package main

import (
	"fmt"
	"slices"
)

// slice is dynamic
// most used
// ++ methods

func main() {
	var nums []int // slice of int
	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)

	// slice literal
	names := []string{"John", "Jane", "Doe"}
	names = append(names, "Alice")

	// slice iteration
	for i, v := range nums {
		println(i, v)
	}

	// slice length
	println(len(nums))

	var num2 = make([]int, 5, 10)

	fmt.Println(len(num2))
	fmt.Println(cap(num2))

	// copy slice
	var num3 = make([]int, 5, 10)
	copy(num3, num2)
	fmt.Println(num3)

	// slice operator
	var num4 = []int{1, 2, 3, 4, 5}
	fmt.Println(num4[2:3])

	fmt.Println(slices.Equal(num2, num3))

	num3 = append(num3, 100)
	fmt.Println(slices.Equal(num2, num3))
}
