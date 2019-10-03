package main

import (
	"errors"
	"fmt"
)

func main() {
	fooErr := foo()
	barErr := errors.Unwrap(fooErr)
	mooErr := errors.Unwrap(barErr)
	catErr := errors.Unwrap(mooErr)
	baseErr := errors.Unwrap(catErr)
	fmt.Printf("fooErr\t%s\n", fooErr)
	fmt.Printf("barErr\t%s\n", barErr)
	fmt.Printf("mooErr\t%s\n", mooErr)
	fmt.Printf("catErr\t%s\n", catErr)
	fmt.Printf("baseErr\t%s\n", baseErr)

}

func foo() error {
	return fmt.Errorf("this error is from FOO - %w", bar())
}

func bar() error {
	return fmt.Errorf("this error is from BAR - %w", moo())
}

func moo() error {
	return fmt.Errorf("this error is from MOO - %w", cat())
}

func cat() error {
	return errors.New("this error is from CAT")
}
