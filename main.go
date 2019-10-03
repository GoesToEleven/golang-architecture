package main

//go:noinline
func foo() int {
	x := 42
	return x
}

//go:noinline
func bar() *int {
	y := 43
	return &y
}

func main() {
	_ = foo()
	_ = bar()
}
