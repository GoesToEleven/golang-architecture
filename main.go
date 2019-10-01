package main

import (
	"errors"
	"fmt"
	"net"
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
	return "", ErrFile{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

func processFile(filename string) error {
	_, err := openFile(filename)
	if err != nil {
		return fmt.Errorf("Error while opening file: %w", err)
	}
	// Do work on stuff
	return nil
}

func main() {
	err := processFile("test.txt")
	if err != nil {
		var fErr ErrFile
		if errors.As(err, &fErr) {
			fmt.Printf("Was unable to do something with file %s\n", fErr.Filename)
		}
		var netErr net.Error
		if errors.As(err, &netErr) {
			if netErr.Temporary() {
				// Retry
			}
		}
		// Some other error
		fmt.Println(err)
	}
}
