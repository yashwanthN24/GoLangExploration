![alt text](image.png)

if we dont wait for the go routine to finish executing, the main function will exit and the go routine will be terminated before it can complete its execution.


because go keyword just assures that the function will be executed in a separate go routine, it does not wait for the function to finish executing before moving on to the next line of code.


![alt text](image-1.png)

channel is just a shared memory space that can be given to all go routines to communicate via memory space. we can use channel to send and receive data between go routines.

they work in FIFO manner, (just like queues)

![alt text](image-2.png)

![alt text](image-3.png)

![alt text](image-4.png)

![alt text](image-5.png)

main goroutine will read from the channel amd process it

![alt text](image-6.png)

when done we close the channel to signal that no more data will be sent on the channel. this is important because if we dont close the channel, the main goroutine will keep waiting for data to be sent on the channel and will never exit.

![alt text](image-7.png)

then all workers will exit and main goroutine will exit as well.

https://chatgpt.com/c/6ae7909e-1b94-83e8-9213-bd7f9054a152