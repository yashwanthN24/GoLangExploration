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

func changeSlice(nums *[]int){
	*nums = []int{0}
	fmt.Println(nums)
	fmt.Println("Address in changeSlice : ", nums)
}


func changeMap(mp *map[string]int){
	fmt.Println(mp , &mp)
	*mp = map[string]int{"price": 40}
	fmt.Println(mp , &mp)
}

func main(){
	num := 1

	fmt.Println("In main num = " , num)
	
	changeNum(&num)
	
	fmt.Println("In main num = " , num)


	fmt.Println(" printing memory address " , &num) // printing address of the variable num 

	nums := []int{1 , 2 , 3};
	fmt.Println("Address of nums in main : " , &nums)
	changeSlice(&nums)

	fmt.Println(nums)
	fmt.Println(&nums)

	myMap := map[string]int{"yash":12}
	
	fmt.Println(myMap , &myMap)
	changeMap(&myMap)
	fmt.Println(myMap , &myMap)
	
}