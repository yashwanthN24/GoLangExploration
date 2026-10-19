package main

import "fmt"

// func printSlice(items []int){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }

// instead of any if writtedn as interface{} also same i.e

// func printSlice[T interface{}](items []T){}

// func printSlice[T interface{}](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T any](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }

func printStringSlice(items []string){
	for _ , item := range items{
		fmt.Println(item)
	}
}

// this makes sure that we are mor stricter to keep our slices to either int slice or string slice only 
// func printSlice[T int | string](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }


// func printSlice[T int | string | bool | float32](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }


// comparable includes all types 
// func printSlice[T comparable](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }
func printSlice[T comparable , V string](items []T , name V){
	for _ , item := range items{
		fmt.Println(item , name)
	}
}


type stack[T any] struct{
	elements []T
}


func main(){

	// nums := []int{1 , 2 , 3}
	// printSlice(nums)

	names := []string{"golang" , "c" , "cpp"}
	// printStringSlice(names)
	// printSlice([]int{1 , 2 , 3 , 5})

	printSlice(names , "yash")

	printSlice([]int{1 , 2 , 3} , "yash")
	printSlice([]bool{false , true ,  true} , "yash")

	myStack := stack[int]{
		elements: []int{1 , 2 , 3},
	}

	fmt.Println(myStack)

	myStack2 := stack[string]{
		elements: []string{"golang" , "cpp"},
	}

	fmt.Println(myStack2)

}