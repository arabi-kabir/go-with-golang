package main

import "fmt"

// enumerated types
type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Shipped   OrderStatus = "Shipped"
	Delivered OrderStatus = "Delivered"
)

type OrderId int

const (
	OrderID1 OrderId = iota
	OrderID2
	OrderID3
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {

	changeOrderStatus(Shipped)

}
