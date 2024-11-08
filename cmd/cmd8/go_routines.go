package main

import (
	"fmt"
	"time"
)

func printnos() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond) // pause for half sec
	}
}
func main() {
	go printnos()               // go routine
	fmt.Print("Hello there \n") // this statement will be printed first, because the program doesnt wait for the func to compleletely get executed.
	time.Sleep(3 * time.Second) // we put a pause, so tht the func finishes executing, without this pause the program will end without printing the func as it is still being processed
}

// to start a goroutine we use the "go" keyword before a func call.

// go routines launch multiple functions and execute them concurrently
// concurrency means multiple tasks are in progress at the same time
// allow programs to perform multiple tasks at the same time without waiting for each task to finish one by one.
// in concurrency if one task takes 3secs to execute, while tht task is in progress, the cpu starts executing other tasks.
// cpu goes back and forth executing tasks.

// in parallet execution, usually there are more than 1 cpu core tht are executing tasks, this also achieves concurrency. here two cpus are working in parallel to execute tasks.
// parallel execution basically executes the task at the same time, there is no back and forth.
// overall concurrency != parallel execution.

// goroutines are lighter than threads, they use less memory

// we make a function a goroutine such tht it achieves concurrency, and the other tasks are peformed while the fuction is still in process.
