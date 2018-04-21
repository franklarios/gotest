package main

import (
	"fmt"

	"github.com/franklarios/gotest/007-level-twelve/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
