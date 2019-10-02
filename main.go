package main

import (
	"io"
	"os"
)

type writeFile struct {
	f   *os.File
	err error
}

func newWriteFile(filename string) *writeFile {
	f, err := os.Create(filename)
	return &writeFile{
		f:   f,
		err: err,
	}
}

func (w *writeFile) WriteString(text string) {
	if w.err != nil {
		return
	}

	_, err := io.WriteString(w.f, text)
	if err != nil {
		w.err = err
	}
}

func (w *writeFile) Close() {
	if w.err != nil {
		return
	}

	err := w.f.Close()
	if err != nil {
		w.err = err
	}
}

func (w *writeFile) Err() error {
	return w.err
}

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(f, "Hello World")
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(f, "More Text!")
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
