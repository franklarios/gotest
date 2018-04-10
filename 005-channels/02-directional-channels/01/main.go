package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2) // buffered channel of type int

	c <- 42 // adds values to channel
	c <- 43

	fmt.Println(<-c) // receives values from channel and prints them
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c) //shows the type of variable
}
