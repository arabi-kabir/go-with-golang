package main

import "fmt"

func main() {

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// classic for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("index", i)
	}

	// use break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}

		fmt.Println("break - index", i)
	}

	// use continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("continue - index", i)
	}

	// range
	for i := range 3 {
		fmt.Println("range3 - index", i)
	}

	// range
	for i := range []int{3, 32, 33, 43, 35} {
		fmt.Println("range - index", i)
	}

	// range with value
	for i, v := range []int{11, 22, 31, 41, 51} {
		fmt.Println("range - index", i, "value", v)
	}
}
