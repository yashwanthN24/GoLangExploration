package main

import "fmt"

func changeSlice(array *[]int){
	fmt.Println(array)
	*array = []int{0}

	fmt.Println(array)
}

func main() {

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	changeSlice(&slice1)
	fmt.Println(slice1)


	fmt.Println()


	fmt.Printf("Underlying array address: %p\n", slice1)

	fmt.Printf("Slice header variable address: %p\n", &slice1)
}


// Sprintf (to print string and store the string inariable)

// to read ?input from user use Scan() passing pointer to variable to store the input same c style input reading like scanf