package main

import "fmt"

func main() {
	var nums [4]int

	nums[0] = 10
	nums[1] = 20
	// nums[2] = 30
	nums[3] = 40

	fmt.Println(nums)

	// array literal
	names := [4]string{"John", "Jane", "Doe"}
	names[3] = "Alice"
	fmt.Println(names)
	fmt.Println(names[2])

	// array iteration
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// array length
	fmt.Println(len(nums))

	var okay [4]bool
	okay[1] = true
	fmt.Println(okay)

	// short
	cities := [...]string{"New York", "Los Angeles", "Chicago"}
	fmt.Println(cities)
	fmt.Println(len(cities))

	country := [3]string{"USA", "Canada", "Mexico"}
	fmt.Println(country)

	// 2D array
	matrix := [2][3]int{{}, {}}
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println(matrix)
}
