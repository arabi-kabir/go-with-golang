package main

import "fmt"

func main() {
	const pi float64 = 3.14

	fmt.Println("Value of pi:", pi)

	const (
		port = 9000
		host = "localhost"
	)

	fmt.Println("Server running on", host, "at port", port)
}
