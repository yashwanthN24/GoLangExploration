![alt text](image.png)

if we dont wait for the go routine to finish executing, the main function will exit and the go routine will be terminated before it can complete its execution.


because go keyword just assures that the function will be executed in a separate go routine, it does not wait for the function to finish executing before moving on to the next line of code.