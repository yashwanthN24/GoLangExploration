package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()
	// this is guranteed to run after this function completes irrespective of any error or normal execution defer enoreces it to run this after this fuction completes so think of it like last line to this function always runs before this function completes

	time.Sleep(time.Millisecond * 50)
	// to add 50 milliseconds delay

	// simulating image processing
	fmt.Printf("image processed : %s\n", url)

	resultChan <- Result{
		Value: url,
		Err:   nil,
	} // sending the data to the channle for the mainroutine to recieve it and process it  (sending to channle )
}

type Result struct {
	Value string
	Err   error
}

func main() {
	var wg sync.WaitGroup
	resultchan := make(chan Result, 6)

	// fmt.Println("welcome to go concurrency")
	startTime := time.Now()
	// fmt.Print(startTime)

	wg.Add(6) // 3 since we are running 3 go routines  pr 3 workers
	go worker("image_1.png", &wg, resultchan)
	go worker("image_2.png", &wg, resultchan)
	go worker("image_3.png", &wg, resultchan)
	go worker("image_4.png", &wg, resultchan)
	go worker("image_5.png", &wg, resultchan)
	go worker("image_6.png", &wg, resultchan)

	// after ach goroutine completion it will decrement the counter pased to add of WaitGroup

	// runs synchronously one after another
	// so total 150ms

	// go func ()  {
	wg.Wait()
	close(resultchan) //close the channel

	// }()

	// aggragting results fan in
	for result := range resultchan {
		fmt.Printf("Recieved : %v \n", result)
		if result.Err != nil{
			// reporcess 
		}
	}

	fmt.Printf("it took %s ms. ", time.Since(startTime))

}

/*

This is fan out pattern of concurrency wherein we are distributing our work to multiple goroutines

you see that now all those 5 workerexecite in 50ms because they run at the same time concyreent via goroutines light weight threads that go intelligently maps to actual os thread concurrently

*/

/*


But how to get data from the goroutines onto the mainthread or main goroutines , channels

channels allow communication between goroutines to send and recieve data

we must define the data being passed in the channel , string int float struct (in production we use struct in real world backend as struct define custom types )

Two types of channels

1) buffered channels (the queue we saw in diagram): we must pass the size of the queue

suppose 5 we csan only add 5 message from the gorouties and after which sender is blocked as channelis full

unbuffered channel boht sides sender and reciever must be ready (connection type)

with goruntines we can do cpu intensive tasks concurrently

*/

// waitgroups , channels and goroutines are actually used in go lang backend for concurrency n golang

// in go we dont thrwo errors rather we return errors from functions
