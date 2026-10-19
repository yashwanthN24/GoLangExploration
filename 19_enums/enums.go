package main

import "fmt"

// unlike other languages enums in go are made using type and const
type Mytype string 

// type OrderStatus int 

// const (
// 	Recieved OrderStatus = iota
// 	Confirmed 
// 	Prepared 
// 	Delivered
// )


type OrderStatus string

const (
	Recieved OrderStatus = "recieved"
	Confirmed OrderStatus = "confirmed"
	Prepared OrderStatus = "prepared"
	Delivered  OrderStatus= "delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to" , status)
}

// enumerated  types 
func main(){
	changeOrderStatus(Delivered)
}