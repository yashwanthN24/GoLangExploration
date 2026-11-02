package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()
	// this is guranteed to run after this function completes irrespective of any error or normal execution defer enoreces it to run this after this fuction completes so think of it like last line to this function always runs before this function completes

	// aggragting results fan in
	for job := range jobsChan {
		// to add 50 milliseconds delay
		time.Sleep(time.Millisecond * 50)
		// simulating image processing
		// fmt.Printf("image processed : %s\n", job)
		resultChan <- Result{
			Value: job,
			Err:   nil,
		} // sending the data to the channle for the mainroutine to recieve it and process it  (sending to channle )
	}

	fmt.Printf("worker shutting down \n")

}

type Result struct {
	Value string
	Err   error
}

func main() {

	jobs := []string{
		"image_1.png",
		"image_2.png",
		"image_3.png",
		"image_4.png",
		"image_5.png",
		"image_6.png",
		"image_7.png",
		"image_8.png",
		"image_9.png",
		"image_10.png",
		"image_11.png",
		"image_2.png",
		"image_12.png",
		"image_13.png",
		"image_14.png",
		"image_15.png",
		"image_16.png",
		"image_17.png",
		"image_18.png",
		"image_19.png",
		"image_20.png",
	}

	var wg sync.WaitGroup
	totalWorkers := 5

	resultchan := make(chan Result, 50)
	jobsChan := make(chan string, len(jobs))

	// fmt.Println("welcome to go concurrency")
	startTime := time.Now()
	// fmt.Print(startTime)

	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultchan) // warming the workers to wait to process that tasks where they wait for jobs in jobschanel

	}

	//works but for 1000 job we are creatuing 1000 goroutines which can degrade the ystem performance so we have to limit the nujmber of goroutines created


	go func() {
		wg.Wait()
		close(resultchan) //close the channel

	}()

	// send the jobs
	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}

	close(jobsChan)

	for result := range resultchan{
		fmt.Printf("` Job completed %v\n" , result)
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
