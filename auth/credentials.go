package auth

import "fmt"

// func loginWithCredentials(username string , password string){
// 	fmt.Println(username , password)
// }

// by default in go  all function and thing that are staring with lowercase letter are availableonly in that package (folder) in which they are defined

// to make them available for other package startnam,eig first letter a Capital letter


func LoginWithCredentials(username string , password string){
	fmt.Println(username , password)
}