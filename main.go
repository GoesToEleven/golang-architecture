package main

import (
	"context"
	"fmt"

	"github.com/GoesToEleven/golang-architecture/session"
)

type key int

var userKey key = 0

func main() {
	ctx := context.Background()
	ctx = session.WithUserID(ctx, 1)

	userID := session.GetUserID(ctx)
	if userID == nil {
		fmt.Println("Not Logged in!")
		return
	}
	fmt.Println(*userID)

	uId := ctx.Value(userKey)
	fmt.Println(uId)
}
