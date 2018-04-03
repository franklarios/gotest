package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Frank",
	}

	saySomething(&p1)

	//saySomething(p1)
	//does not work because it is of type person and not a pointer

	p1.speak() //This works however
}
