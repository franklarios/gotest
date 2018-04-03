package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Beginning CPUs:", runtime.NumCPU())
	fmt.Println("Beginning Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()

	fmt.Println("Mid # CPUs:", runtime.NumCPU())
	fmt.Println("Mid # Goroutines:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("Ending CPUs:", runtime.NumCPU())
	fmt.Println("Ending Goroutines:", runtime.NumGoroutine())
}
