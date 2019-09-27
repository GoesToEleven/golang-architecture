package main

import (
	"fmt"

	"github.com/GoesToEleven/golang-architecture"
	"github.com/GoesToEleven/golang-architecture/storage/harddrive"
)

func main() {
	storage := harddrive.Db{}

	u1 := architecture.User{
		First: "James",
	}

	u2 := architecture.User{
		First: "Jenny",
	}

	architecture.Put(storage, 1, u1)
	architecture.Put(storage, 2, u2)

	fmt.Println(architecture.Get(storage, 1))
	fmt.Println(architecture.Get(storage, 2))
}
