package main

import(
	"fmt"
	"github.com/franklarios/gotest/007-testing/04-benchmark/00-live/saying"

)

func main() {
	fmt.Println(saying.Greet("James"))
}