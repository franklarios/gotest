package main

import (
	"fmt"
)

func main() {
	c := make(<-chan int, 2) //"<-" before "chan" indicates this is a receive-only type channel

	c <- 42 // operation invalid because you cannot send to receive-only type channel
	c <- 43

	fmt.Println(<-c) 
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
