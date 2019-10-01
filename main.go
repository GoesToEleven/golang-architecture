package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)
	ctx = context.WithValue(ctx, 1, "admin")
	if v := ctx.Value("userID"); v != nil {
		fmt.Println(v)
	} else {
		fmt.Println("no value associated with that key")
	}

	if v := ctx.Value(1); v != nil {
		fmt.Println(v)
	} else {
		fmt.Println("no value associated with that key")
	}

	if v := ctx.Value(2); v != nil {
		fmt.Println(v)
	} else {
		fmt.Println("no value associated with that key")
	}
}
