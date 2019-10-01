package main

import (
	"context"
	"fmt"

	"github.com/GoesToEleven/golang-architecture/session"
)

func main() {
	ctx := context.Background()
	ctx = session.SetUserID(ctx, 1)
	ctx = session.SetAdminAccess(ctx, true)

	uID := session.GetUserID(ctx)
	isAdmin := session.GetAdmin(ctx)
	fmt.Printf("User %d is an admin %t\n", uID, isAdmin)
}
