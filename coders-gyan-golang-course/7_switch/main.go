package main

import (
	"fmt"
	"time"
)

func main() {

	const age int = 30

	// simple switch
	switch age {
	case 30:
		fmt.Println("You are 30 years old")
	case 40:
		fmt.Println("You are 40 years old")
	default:
		fmt.Println("Age is not 30 or 40")
	}

	// multiple switch cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whoAMI := func(i interface{}) string {
		switch v := i.(type) {
		case string:
			return "Hello, " + v
		case int:
			return fmt.Sprintf("You are %d years old", v)
		default:
			return "I don't know who you are"
		}
	}

	fmt.Println(whoAMI("John"))
	fmt.Println(whoAMI(55))
}
