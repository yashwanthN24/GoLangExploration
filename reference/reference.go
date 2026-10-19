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