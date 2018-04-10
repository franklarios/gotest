package main

import (
	"fmt"
)

func main() {
	c := make(chan<- int, 2) //"<-" after "chan" indicates this is send-only type channel

	c <- 42 //sends value to channel
	c <- 43

	fmt.Println(<-c) // operation invalid because you cannot receive (pull values) from send-only type channel
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
