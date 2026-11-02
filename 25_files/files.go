package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		//   log the error
		panic(err)
	}

	fmt.Println("file name : " , fileInfo.Name())
	fmt.Println("Is directory : " , fileInfo.IsDir())
	fmt.Println("file size : " , fileInfo.Size()) // number oif character i.e bytes each character 1 byte 
	fmt.Println("file permission : " , fileInfo.Mode())
	fmt.Println("file modified at : " , fileInfo.ModTime())

}
