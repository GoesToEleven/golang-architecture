package session

import (
	"context"
)

type key int

var userKey key = 0

func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userKey, userID)
}

func GetUserID(ctx context.Context) *int {
	userID, ok := ctx.Value(userKey).(int)
	if !ok {
		return nil
	}
	return &userID
}
