package main

import "fmt"

// for -> only construct in go for looping
func main(){
	// only for loop is used in go for all looping 

	// while loop 
	i := 1 
	for i <= 3 {
		fmt.Println(i)
		// i = i+1
		i += 1
	}

	// infinite loop 
	// for {
	// 	fmt.Println(1)
	// }


	// classic for loop 
	// for i:=1; i<=3 ; i++{
	// 	fmt.Println(i)
	// }

	// break and continue 
	fmt.Println()
	for i:=1 ; i<=10; i++{
		if(i == 5){
			break
		}
		fmt.Println(i)
	}

	fmt.Println()

	for i:=1; i<=8;i++{
		if (i == 5){
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// range
	for i:= range 3{
		fmt.Println(i)
	} 


	fmt.Println()


	for i:= range 10{
		fmt.Print(i , " ")
	}

	fmt.Println()

	//  prints 0 to 10 ( 11 values)
	for i:=range 11{
		fmt.Println(i)
	}
}