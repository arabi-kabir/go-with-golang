package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	email string
	age   int
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  customer // struct embedding
}

// constructor function
func newOrder(id string, amount float32, status string) *order {
	return &order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// receiver function
func (o *order) changeStatus(status string) {
	o.status = status
}

//

func main() {
	order1 := order{
		id:     "123",
		amount: 100.50,
		status: "pending",
	}

	order1.createdAt = time.Now()

	order1.changeStatus("confirmed")

	fmt.Println(order1)

	fmt.Println("Order ID:", order1.id)

	customer1 := customer{
		name:  "John Doe",
		email: "test@gmail.com",
		age:   30,
	}

	order1.customer = customer1

	fmt.Println(order1)

	//
	order2 := newOrder("456", 200.75, "pending")
	fmt.Println(order2)

	order2.changeStatus("shipped")
	fmt.Println(order2)

	langiage := struct {
		name    string
		version float32
	}{"Go", 1.20}

	fmt.Println(langiage)
}
