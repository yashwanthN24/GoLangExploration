package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	// this is guranteed to run after this function completes irrespective of any error or normal execution defer enoreces it to run this after this fuction completes so think of it like last line to this function always runs before this function completes

	time.Sleep(time.Millisecond * 50)
	// to add 50 milliseconds delay

	// simulating image processing
	fmt.Printf("image processed : %s\n", url)

}

func main() {
	var wg sync.WaitGroup
	// fmt.Println("welcome to go concurrency")
	startTime := time.Now()
	// fmt.Print(startTime)

	wg.Add(6) // 3 since we are running 3 go routines  pr 3 workers
	go worker("image_1.png", &wg)
	go worker("image_2.png", &wg)
	go worker("image_3.png", &wg)
	go worker("image_4.png", &wg)
	go worker("image_5.png", &wg)
	go worker("image_6.png", &wg)

	// after ach goroutine completion it will decrement the counter pased to add of WaitGroup

	// runs synchronously one after another
	// so total 150ms

	wg.Wait()
	fmt.Printf("it took %s ms. ", time.Since(startTime))

}

/*

This is fan out pattern of concurrency wherein we are distributing our work to multiple goroutines

you see that now all those 5 workerexecite in 50ms because they run at the same time concyreent via goroutines light weight threads that go intelligently maps to actual os thread concurrently

*/

/*


But how to get data from the goroutines onto the mainthread or main goroutines , channels 

channels allow communication between goroutines to send and recieve data

*/
