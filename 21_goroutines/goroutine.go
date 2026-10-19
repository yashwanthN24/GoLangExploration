// goroutine are go's loght weight thread that handle task parallely using efficient threads switch for io and waiting  and resume other oejdning goroutines

package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing task " , id);
}

func main(){
	for i:=0 ; i<=10 ; i++{
		// go task(i)

		go func (i int)  {
			fmt.Println(i)
		}(i)
	}


	time.Sleep(time.Second*1)
}

// go routine is concurreny of threads multithreading 
// ligjtweight thread for io task they are paused and let other thread to execute this was concurreny is achieved with multiple lightweight threads hence fast 