package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//   log the error
	// 	panic(err)
	// }

	// fmt.Println("file name : " , fileInfo.Name())
	// fmt.Println("Is directory : " , fileInfo.IsDir())
	// fmt.Println("file size : " , fileInfo.Size()) // number oif character i.e bytes each character 1 byte
	// fmt.Println("file permission : " , fileInfo.Mode())
	// fmt.Println("file modified at : " , fileInfo.ModTime())

	//? Reading file

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // now need to store the content of this file in a bufffer or in memory that i s available as long as this prgram (process exists /runs)

	// // buf := make([]byte, 12)
	// fileInfo, err := f.Stat()
	// buf := make([]byte, fileInfo.Size())

	// d, err := f.Read(buf) // d is the number of bytes read (bytes means characters as 1 character is 1 byte in memory )

	// if err != nil {
	// 	panic(err)
	// }

	// println("data", d, buf)

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	//? Reading file easy way

	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// Readfile is good for small files as it loads the entire file nto memory at once
	// so for large files very slow so we have to use treaming like nodejs instead of fs in nodejs

	//? reading folders

	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)// if number < 0 it reads all contents of he directory
	// if n > 0 it gets only those manay number sof directory or files as a slice

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//? creating a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi Go\n")
	// f.WriteString("Nice language\n")

	// bytes := []byte("Hello GoLang")

	// f.Write(bytes) // overides the contents in file other way to write to files via bytes slice

	// appends to existing content

	// then how to overide to override use os.writeFile

	//? Read and write to other file (streaming fashion For larghe file sproduction style)

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile) // this already ha a buffer of default size 4096
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println(" written to new file successfully ")

	//? deleting a file

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted sucessfully")
}
