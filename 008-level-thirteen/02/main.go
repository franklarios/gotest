package main

import (
	"fmt"

	"github.com/franklarios/gotest/008-level-thirteen/02/quote"
	"github.com/franklarios/gotest/008-level-thirteen/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
