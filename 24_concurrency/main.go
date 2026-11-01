package main

import (
	"fmt"
	"time"
)

func worker(url string) string {
	time.Sleep(time.Millisecond * 50)
	// to add 50 milliseconds delay

	// simulating image processing
	fmt.Printf("image processed : %s\n", url)
	return url

}

func main() {
	// fmt.Println("welcome to go concurrency")
	startTime := time.Now()
	// fmt.Print(startTime)

	result := worker("image_1.png")
	result2 := worker("image_2.png")
	result3 := worker("image_3.png")
	fmt.Println("resulT:", result2)
	fmt.Println("resulT:", result3)
	fmt.Println("resulT:", result)

	// runs synchronously one after another 
	// so total 150ms

	fmt.Printf("it took %s ms. ", time.Since(startTime))

}
