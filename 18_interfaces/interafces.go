package main

import "fmt"

// interface are contact that custom classes can implement
// convention is to use ter for interfaces

type paymenter interface{
	pay(amount float32)
	refund(amount float32 , account string)
}


type payment struct{
	gateway paymenter
}

// Open close principle open for extension but closed fomodufication 
func (p payment) makePayment(amount float32){
	// razorpayPaymentGW := razorpay{}
	// stripePaymentGw := stripe{}
	// // razorpayPaymentGW.pay(amount)
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

// unlike other language you dont need to implement a interface in any struct if you create same method signature as that interface method then it automatically implement it this is go power 


type razorpay struct{}

func (r razorpay) pay(amount float32){
	// logic to make payment 
	fmt.Println("making payment using razorpay" ,amount)
}


type stripe struct{}

func (s stripe) pay(amount float32){
	// logic to make payment 
	fmt.Println("making payment using stripe" ,amount)
}


type payPal struct{}

func (p payPal) pay(amount float32){
	fmt.Println("making payment using payPal " , amount)
}

func (p payPal) refund(amount float32 , account string){

}

func main(){
	// razorpayGW := razorpay{}
	// newPayment := payment{
	// 	gateway: razorpayGW,
	// }
	// newPayment.makePayment(100)


	paypalGw := payPal{}

	newPayment := payment{
		gateway:paypalGw,
	}
	newPayment.makePayment(1000)
}