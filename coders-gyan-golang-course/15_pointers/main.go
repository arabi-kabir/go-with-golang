package main

import "fmt"

func changeNumber(num int) {
	num = 20
	fmt.Println("Inside changeNumber function:", num)
}

func changeNumberPointer(num *int) {
	*num = 20
	fmt.Println("Inside changeNumberPointer function:", *num)
}

func main() {
	num := 1

	changeNumber(num)
	println("After changeNumber function call:", num)

	println("memory address ", &num)
	changeNumberPointer(&num)
	println("After changeNumberPointer function call:", num)
}
