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
	f := newWriteFile("file.txt")
	f.WriteString("Hello World")
	f.WriteString("More Text!")
	f.Close()

	err := f.Err()
	if err != nil {
		panic(err)
	}
}
