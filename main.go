package main

import (
	"fmt"
	"log"
)

type errorString string

func (es errorString) Error() string {
	return fmt.Sprintf("this is an error string with info about who what where when why how - %s", string(es))
}

func main() {
	n, err := sum(2, 4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

func sum(i, j int) (int, error) {
	n := i + j
	if n != i+j {
		var sErr errorString = "the numbers didn't add up"
		return 0, sErr
	}
	return n, nil
}
