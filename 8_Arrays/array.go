package main

import "fmt"


func main(){
	// intilaized to default values i.e 0 for int 
	// int -> 0 , string -> "" , bool -> false 
	// var nums[4] int 

	// fmt.Print(len(nums))

	// fmt.Println("\n" , nums)

	// nums[0] = 44
	// fmt.Println(nums[0])

	// var visted[10] bool; // instialized to default value false 

	// visted[9] = true
	// fmt.Println(visted)

	var names[10] string // array of empty strings 
	fmt.Println(names)

	fmt.Println(len(names))

	// to declare it in a single line and intialize 
	nums := [3]int{1 , 2 , 3}
	fmt.Println(nums)

	// 2D Arrays  
	nums2 := [2][2]int{{3, 4} , {5 , 6}}
	fmt.Println(nums2)

	// Use array when 
	// - Fixed size , that is predictable
	// - Memory optimization 
	// - Constant time access 

	// generally in go we use slices more  than arrays 
}