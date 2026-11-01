package main

import (
	"fmt"
	"time"
)

func worker(url string) {
	time.Sleep(time.Millisecond * 50)
	// to add 50 milliseconds delay

	// simulating image processing
	fmt.Printf("image processed : %s\n", url)

}

func main() {
	// fmt.Println("welcome to go concurrency")
	startTime := time.Now()
	// fmt.Print(startTime)

	go worker("image_1.png")
	go worker("image_2.png")
	go worker("image_3.png")

	// runs synchronously one after another
	// so total 150ms

	fmt.Printf("it took %s ms. ", time.Since(startTime))

}
