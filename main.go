package main

import (
	"context"
	"fmt"
)

type key int

var userKey key = 0
var ipKey key = 1
var isAdminKey key = 2
var sessionKey key = 3

func foo(ctx context.Context) {
	ctx = context.WithValue(ctx, "currentFunc", "foo")


}

func main() {
	ctx := context.WithValue(context.Background(), userKey, 1)
	// ctx := context.Background()

	userID, ok := ctx.Value(userKey).(int)
	if !ok {
		fmt.Println("Not Logged in!")
		return
	}
	fmt.Println(userID)
}
