package main

import (
	"fmt"
	"errors"
)

func cat() error {
	return errors.New("cat is an error")
}

func moo() error {
	return fmt.Errorf("moo is an error: %w", cat())
}

func bar() error {
	return fmt.Errorf("bar is an error: %w", moo())
}

func foo() error {
	return fmt.Errorf("foo is an error: %w", bar())
}

func main() {
	err := foo()
	fmt.Println(err)

	baseErr := errors.Unwrap(err)
	fmt.Println(baseErr)

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)
}
