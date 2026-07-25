package main

func main() {

	age := 20

	if age < 18 {
		println("You are a minor")
	} else if age >= 18 && age < 65 {
		println("You are an adult")
	} else {
		println("You are a senior citizen")
	}

	var role string = "admin"
	var hasPermission bool = true

	if role == "admin" && hasPermission {
		println("You have admin access")
	} else if role == "user" && hasPermission {
		println("You have user access")
	} else {
		println("Access denied")
	}

	// assignment in if statement
	if score := 85; score >= 90 {
		println("Excellent")
	} else if score >= 75 {
		println("Good")
	} else {
		println("Needs Improvement")
	}

}
