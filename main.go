package main

import "fmt"

type hotdog int

func (h hotdog) cook() {
	fmt.Println("I am cooking")
}

type hotFood interface {
	cook()
}

func bar(hf hotFood) {
	hf.cook()
}

func main() {
	var x hotdog = 42
	var y hotFood
	y = x
	fmt.Printf("%T\n", y)

	bar(x)
}
