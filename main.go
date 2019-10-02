package main

import (
	"github.com/GoesToEleven/golang-architecture/fileWriter"
)

func main() {
	f := fileWriter.NewWriteFile("file.txt")
	f.WriteString("Hello World")
	f.WriteString("More Text!")
	f.Close()

	err := f.Err()
	if err != nil {
		panic(err)
	}
}
