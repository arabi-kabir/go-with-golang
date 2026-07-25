package main

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// func sumAny(nums ...interface{}) int {
// 	total := 0
// 	for _, v := range nums {
// 		switch v.(type) {
// 		case int:
// 			total += v.(int)
// 		case float64:
// 			total += int(v.(float64))
// 		default:
// 			println("Unsupported type")
// 		}
// 	}
// 	return total
// }

func main() {

	println(sum(1, 2, 3, 4, 5))

	myArr := []int{1, 2, 3, 4, 5}

	println(sum(myArr...))

}
