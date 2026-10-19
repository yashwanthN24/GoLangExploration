package main

import "fmt"

// func sum(vals ...int){
// 	total := 0

// 	for _ , num:= range vals{
// 		total += num
// 	}

// 	fmt.Println(total)
// }


func sum(vals ...int)int{
	total := 0 

	for _ , num:= range vals{
		total += num
	} 

	return  total
}


// interface{} is like any
func customPrint(values ...interface{}){
	fmt.Println(values...) // spreads the entire slice that we received
}



func main(){


	// This println can accept varibalke lenegth paratmetrs how variadic function like varargs in java and rest operator in js 
	fmt.Println(1 , 2 , 3 , 4 , 5 , 6 , "hello")
	fmt.Println(sum(1 , 2 , 3 , 4 , 5 , 6))

	customPrint(1 , "hello" , true, false)

	numbers := []int{12 , 4 , 6}
	fmt.Println(sum(numbers...)) // spreads the values of the numbers slice and pased them to function
}