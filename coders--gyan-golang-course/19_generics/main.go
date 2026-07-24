package main

import "fmt"

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSliceOnly[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type Stack[T any] struct {
	elements []T
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}
	printSlice(nums)
	printSlice(names)

	printSliceOnly(nums)
	printSliceOnly(names)

	myStack := Stack[int]{
		elements: []int{1, 2, 3},
	}

	myStack1 := Stack[string]{
		elements: []string{"a", "b", "c"},
	}

	fmt.Println("Stack elements:", myStack.elements)
	fmt.Println("Stack elements:", myStack1.elements)

}
