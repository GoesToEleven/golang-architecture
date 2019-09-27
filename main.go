package main

import (
	"fmt"
)

type user struct {
	first string
}

type mongo map[int]user

func (m mongo) save(n int, u user) {
	m[n] = u
}

func (m mongo) retrieve(n int) user {
	return m[n]
}

type harddrive map[int]user

func (hd harddrive) save(n int, u user) {
	hd[n] = u
}

func (hd harddrive) retrieve(n int) user {
	return hd[n]
}

type accessor interface {
	save(n int, u user)
	retrieve(n int) user
}

func put(a accessor, n int, u user) {
	a.save(n, u)
}

func get(a accessor, n int) user {
	return a.retrieve(n)
}

func main() {
	storage := harddrive{}

	u1 := user{
		first: "James",
	}

	u2 := user{
		first: "Jenny",
	}

	put(storage, 1, u1)
	put(storage, 2, u2)

	fmt.Println(get(storage, 1))
	fmt.Println(get(storage, 2))
}
