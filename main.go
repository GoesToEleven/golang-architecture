package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("rumi.txt")

	if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("you do not have permission to open the file: %w", err)
		log.Println(err)
	} else if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("the file does not exist: %w", err)
		log.Println(err)
	} else if err != nil {
		err = fmt.Errorf("file couldn't be opened: %w", err)
		log.Println(err)
	}
/*
	if err == os.ErrPermission {
		err = fmt.Errorf("you do not have permission to open the file: %w", err)
		log.Println(err)
	} else if err == os.ErrNotExist {
		err = fmt.Errorf("the file does not exist: %w", err)
		log.Println(err)
	} else if err != nil {
		// panic(err)

		err = fmt.Errorf("file couldn't be opened: %w", err)
		log.Println(err)
	}
*/
	defer f.Close()
}
