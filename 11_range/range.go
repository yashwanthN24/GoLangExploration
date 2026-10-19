package main

import "fmt"

// iterating over data structures
func main(){

	nums := []int{1 , 2 ,3}

	for i:=0; i<len(nums);i++{
		fmt.Println(nums[i]) // prints the value of nums at index i-1
	}

	fmt.Println()


	// range function 

	for _ , num := range nums{
		// first value is index , num is the actual value 
		fmt.Println(num)
	}

	// sum 

	sum := 0 

	for _ , num := range nums{
		sum += num
	}

	fmt.Println(sum)

	fmt.Println()


	for index , num := range nums{
		fmt.Println("index : " , index , "value : " , num)
	}


	// iterating map via range 

	m:= map[string]string{"fname": "john" , "lname": "doe"}

	for k,v := range m{
		fmt.Println(k , v)
	}

	fmt.Println()
	for k := range m{
		fmt.Println(k)
	}

	// only keys 

	// iterating string 

	// unicode code point rune
	
	for _ , ch := range "yash"{
		fmt.Println(ch) // prints the unicode of the character (ascii) //
	}

	fmt.Println()

	// i is starting byte of rune
	//  300 -> 1 byte , 2byte 
	for i , c := range "yash"{
		fmt.Println(i , c)
	}

	fmt.Println()


	for _ , ch := range "yashwanth"{
		fmt.Println(string(ch) , ch)
	}
}