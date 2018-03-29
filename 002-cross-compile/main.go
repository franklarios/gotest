package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime: \t%s\narchitecture: \t%s\n", runtime.GOOS, runtime.GOARCH)
}
