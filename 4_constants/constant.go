package main

import "fmt"

var namqe string = "S"
const nam3237 = 782.4383

// here you cant use := synatx 

// name346663 := "Yashwaja" // error

func main()  {
	const name = "String" // type infered from right value
	const age = 30

	const fracional float64 = 56
	// fracional = 37
	fmt.Println(fracional)

	fmt.Println(namqe)
	fmt.Println(nam3237)

	const (
		port = 5000
		host = "localhost"
	)
	// creating mutiple constants 

	// port  = 4500;

	fmt.Println(port , host)

}