package main

import "fmt"

// same concept as c

// just like js function name gives reference

// value is coped for reference types reference is copied

// so changing this variable in function wont affect the orginal variable it affects onluy the local variable of this function
// func changeNum(num int){
// 	num = 5
// 	fmt.Println(" In Change Num " ,  num)
// }

// by reference via pointers
func changeNum(num *int){
	*num = 5
	fmt.Println(" In Change Num " ,  *num)
	fmt.Println(" In Change Num " ,  num)
}

func main(){
	num := 1

	fmt.Println("In main num = " , num)
	
	changeNum(&num)
	
	fmt.Println("In main num = " , num)


	fmt.Println(" printing memory address " , &num) // printing address of the variable num 

}