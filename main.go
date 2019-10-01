package main

import (
	"errors"
	"fmt"
	"time"
)

type ErrFileNotFound struct {
	Filename string
	When     time.Time
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found at %v", e.Filename, e.When)
}

func (e ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

var ErrNotExist = fmt.Errorf("File does not exist")
var ErrUserNotExist = errors.New("User does not exist")

func main() {
	err := ErrFileNotFound{
		Filename: "test.txt",
		When:     time.Now(),
	}

	fmt.Println(errors.Is(err, ErrFileNotFound{}))

	fmt.Println(err)
}
