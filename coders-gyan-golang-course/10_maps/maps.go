package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {

	// creating map
	m := make(map[string]int)
	m["John"] = 30
	m["Jane"] = 25

	// map literal
	m2 := map[string]int{
		"name": 30,
		"age":  25,
	}

	// map iteration
	for k, v := range m {
		println(k, v)
	}

	// map length
	println(len(m))

	fmt.Println(m2)

	// delete map
	delete(m, "John")

	fmt.Println(m)

	// short map declaration
	age := map[string]int{"John": 30, "Jane": 25}
	fmt.Println(age)

	// check if key exists
	val, ok := age["John"] // value, bool

	fmt.Println(val)

	if ok {
		fmt.Println("John exists")
	} else {
		fmt.Println("John does not exist")
	}

	_, ok2 := age["Doe"]
	if ok2 {
		fmt.Println("Doe exists")
	} else {
		fmt.Println("Doe does not exist")
	}

	// map eqality
	m3 := map[string]int{"John": 30, "Jane": 25}
	m4 := map[string]int{"John": 30, "Jane": 25}
	m5 := map[string]int{"John": 30, "Jane": 26}

	fmt.Println(maps.Equal(m3, m4)) // true
	fmt.Println(maps.Equal(m3, m5)) // false
}
