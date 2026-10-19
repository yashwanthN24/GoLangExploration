package main

import (
	"fmt"
	"time"
)


func main(){

	// simple switch 
	i := 3
	switch i {
		case 1: fmt.Println("one")
		case 2: fmt.Println("two")
		case 3: fmt.Println("three")
		default: fmt.Println("other")
	}


	// multilple condition switch 
	switch time.Now().Weekday() {
	case time.Saturday , time.Sunday :
		fmt.Println(" Its weekend")
	default: 
		fmt.Println("Its working day")
		
	}


	// type switch 

	// functions are first class citizend in go can be assined to varibale and passed as parameter to function and returned from functions 

	whoAmI := func (i interface{}){
		// switch t:= i.(type){
		// case int: 
		// 	fmt.Println("I am int")
		// case string: 	
		// 	fmt.Println("I am String")

		// case bool:
		// 	fmt.Println("I am boolean")

		// default: 
		// 	fmt.Println("Other" , t)
		// }
		switch i.(type){
		case int: 
			fmt.Println("I am int")
		case string: 	
			fmt.Println("I am String")

		case bool:
			fmt.Println("I am boolean")

		default: 
			fmt.Println("Other" )
		}
	} 

	// interface {} : mean's any type

	whoAmI("goLang")

	whoAmI(55.9)
	whoAmI(55)
	whoAmI(true)
	whoAmI(false)
}