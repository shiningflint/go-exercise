package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doJob(i int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d is done!\n", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	fmt.Println(runtime.NumGoroutine()) // Starts from 1 goroutine

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doJob(i, &wg) // Adds 1 goroutine
		fmt.Println(runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine()) // Goes back to 1 goroutine after the wait
}
