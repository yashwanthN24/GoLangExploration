package main

import (
	"fmt"
	"time"
)

// go has no classes

// structs are used to make custom data structures

// order struct

type order struct{
	id string 
	amount float32
	status string 
	createdAt time.Time  // nano second precision very powerful (very accurate )
}

// constructor

// convention is new<structName> 
func newOrder(id string , amount float32 , status string) *order {

	// intial setup goes here......
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// attaching methods on order struct

// reciver is the struct name first letter thats the convention 


// without pointer its a local copy so changing his wont change the actual object so need a pointer
// func (o order) changeStatus(status string){
// 	o.status = status
// }
func (o *order) changeStatus(status string){
	o.status = status // you might ask where is the defrence synatx for struct go handles it for us
}


// dont use pointer when you just want to read the value from the object and dont want to modify it 

// if you want to modify it then use pointers only``
func (o order) getAmount() float32{
	return o.amount
}

func main()  {

	//  if you don't set any field, default value is zero values 
	// int => 0 , float => 0 ,string => "" , bool => false
	var myOrder order = order{
		id: "1",
		amount: 50.00,
		// status: "recieved",
	} 

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder.status)

	

	// structs are like blueprints we c an create multiple instance of it 

	myorder2 := order{
		id: "2",
		amount: 100,
		status: "recieved",
		createdAt: time.Now(),
	}

	myorder2.status = "paid"

	myorder2.changeStatus("yash")

	fmt.Println()
	fmt.Println("Order struct" , myorder2)
	fmt.Println("Order struct" , myOrder)


	myOrder3 := newOrder("3" , 30.35 , "recieved")
	fmt.Println(myOrder3)
	fmt.Println(myOrder3.amount) // dont need to deference the pointer struct handles that 



	// inline struct : when you want to use struct only once (inline) when you  dont want to create multiple instance of structs . use inline structs

	language := struct{
		name string 
		isGood bool
	}{"golang" , true}

	fmt.Println(language)
}
