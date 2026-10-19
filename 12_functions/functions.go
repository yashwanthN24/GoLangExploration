package main

import "fmt"

// func add(a int , b int) int {}
func add(a , b int) int {
	return a+b

	// by default go we have to specify what value we are returning after that () list 
}

// if parameters of the function we can omit the first value type except the last like a , b int 

// go functions can return multiple values 

// we are telling we will return three strings 

func getLanguages() (string , string , string){
	return "golang" , "python" , "java"
}
func getLanguages2() (string , string , bool){
	return "golang" , "python" , false
}

// generally for go project we wil return only two values from a function one the value return by the function and other is the error returned by the function 

// so that we can handle based on whther any error or not 

func processIt(fn func(a int)){
	fn(1)

	// so here fn is of type func()  taking in a int a and returining int
}

func processIt2() func(int) int{
	return func(i int) int {
		return i+1
	}
}

func main(){
	fmt.Println(add(2 , 3))

	result:= add(12 , 3)

	fmt.Println(result)

	fmt.Println(getLanguages())

	lang1 , lang2 , _ :=  getLanguages2()

	// use _ to suppres declared but nor ysed error and when you dont want to use the value returned from the function
	fmt.Println(lang1 , lang2 )



	fn := func (a int)  {
		fmt.Println(2)
	}

	processIt(fn)


	fn2 := processIt2()
	fmt.Println(fn2(10))

}

// go functions are first class citizens they can be assigned to variables and pased to functiona ndretuedn from functions 

// main itself is a function dont need to call it its the entry point of go program 

