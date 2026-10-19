package main

import (
	"fmt"
	"time"
)

// inheritance like in OOPs

type customer struct{
	name string
	phone string
}



// composition
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision (very accurate)
	customer
}

// so now if customer struct has methods that also will be availabe l to order struct 






func main()  {

	newcustomer := customer{
		name: "yashwanth",
		phone: "1828828282",
	}
	newOrder := order{
		id: "1",
		amount: 32,
		status: "recieved",
		customer:newcustomer,
	}

	order:= order{
		id: "2",
		amount: 42,
		status: "recieved",
		customer: customer{
			name: "tara",
			phone: "82822929029",	
		},
	}

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)


	fmt.Println()

	order.customer.name = "robin"

	fmt.Println(order)
	fmt.Println(order.customer)
}

