package main

func add(a int, b int) int {
	return a + b
}

func details(name string, age int) (string, int, string) {
	return name, age, "Engineer"
}

func processIt(fn func(a int, b int) int) {
	// return fn(a, b)
}

func main() {

	println(add(2, 3))

	println(details("John", 30))

	a, b, c := details("Jane", 25)

	println(a, b, c)

}
