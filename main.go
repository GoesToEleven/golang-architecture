package main

import (
	"errors"
	"fmt"
)

type ErrFile struct {
	Filename string
	Base     error
}

func (e ErrFile) Error() string {
	return fmt.Sprintf("File %s: %v", e.Filename, e.Base)
}

func (e ErrFile) Unwrap() error {
	return e.Base
}

var ErrNotExist = fmt.Errorf("File does not exist")

func openFile(filename string) (string, error) {
	return "", ErrNotExist
}

func openFile2(filename string) (string, error) {
	return "", ErrFile{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

func main() {
	_, err := openFile("test.txt")
	if err != nil {
		wrappedErr := fmt.Errorf("Unable to open file %v: %w", "test.txt", err)
		if errors.Is(wrappedErr, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		fmt.Println(wrappedErr)
	}

	_, err = openFile2("test.txt")
	if err != nil {
		if errors.Is(err, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		fmt.Println(err)
	}
}
