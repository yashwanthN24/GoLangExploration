package main

import "fmt"

func main()  {
	// int optimized go optmized whwther to use int8 or int16 etc 
	var number int = 23
	fmt.Println(number)
	// infers the type based on the value on right side
	var num2 = 25
	fmt.Println(num2)

	// Shorthane usefule when value is assigned immediately 
	num3 := 45 
	fmt.Println(num3)

	// suppose i just want to declare a variable and assign value to it afterwards then should do normal way only 
	var text string; 
	text = "Yashwant"
	fmt.Println(text)

	// bool 
	isAdult := true
	fmt.Println(isAdult)

	// floats - float32 , float64
	fraction := 45.2
	fmt.Println(fraction)

	var fraction2 float32 = 34.243
	fmt.Println(fraction2)

	
}