package main

import (
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

	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// now need to store the content of this file in a bufffer or in memory that i s available as long as this prgram (process exists /runs)

	// buf := make([]byte, 12)
	fileInfo, err := f.Stat()
	buf := make([]byte, fileInfo.Size())

	d, err := f.Read(buf) // d is the number of bytes read (bytes means characters as 1 character is 1 byte in memory )

	if err != nil {
		panic(err)
	}

	println("data", d, buf)

	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[i]))
	}
}
