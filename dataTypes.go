package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yash/trial/auth"
	u "github.com/yash/trial/user"
)


func main(){
	// primitive types 

	var isAdult bool = false
	var floating float32 = 34.27
	var name string = "Yashwant"
	var char rune = 'A'
	var complexData complex128 = 127 + 7i
	fmt.Printf("%T and %v \n" , isAdult , isAdult)
	fmt.Printf("%T and %v \n" , floating , floating)
	fmt.Printf("%T and %v \n" , name , name)
	fmt.Printf("%T and %v \n" , char , char)
	fmt.Printf("%T and %v \n" , complexData , complexData)

	// reference types 
	// slices []T , Maps map[k]v , Pointers *T , functions func(){} 

	// for pritimitives default values are zero values i.e for int 0 , for string "" , for bool false

	// for reference types default values are nil 
	
	// slice 
	var costs []int 
	fmt.Printf("%T and %v \n" , costs , costs)
 
	// costs is not set to anthing itailly its nil i.e (null)
	fmt.Println(costs == nil)

	

	fn := func ()  {
		
	}
	fmt.Printf("%T and %v \n" , fn , fn) // so only name of it print reference address of this function at whihc its loacated 

	var maps map[string]int = map[string]int{"price":40}
	fmt.Printf("%T and %v \n" , maps , maps) 

	session := auth.GetSession()
	fmt.Println(session)


	auth.LoginWithCredentials("yash" , "taha")

	user := u.User{
		Name: "yashwanth",
		Email: "yash@220@gamil.com",
	}
	fmt.Println(user)

	color.Green(user.Email)
	
}


/*

Simple rule if you want anything to be available only in your package (foldee) write them as startletter lowercase 

if first letter is captialized then they are avilable globally to other packages or main when imported 

These are our package 

but go themselves has package build and shared like npm packages 


go packages search in google 

find the package 
say i fond this package 

https://pkg.go.dev/github.com/fatih/color


https://pkg.go.dev/github.com/fatih/color

copy the package name at top something like this github.com/username/packageName


install this package using command 

go get github.com/fatih/color

go mod tidy (to fix things say you imported directly a package but not did go get then run this or ay other ussue run this )

go.sum is like package-lock.json finalized lockthe version of your project 


To create a new go project we use 

go mod init github.com/username/reponame 

This is like doing npm init for nodjes project same but for go 

*/