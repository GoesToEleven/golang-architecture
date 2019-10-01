package session

import "context"

type stringKey string
type intKey int

var userID stringKey
var admin intKey

func SetUserId(ctx context.Context, uID int) context.Context {
	return context.WithValue(ctx, userID, uID)
}

func SetIsAdmin(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

func GetUserId(ctx context.Context) int {
	if uid := ctx.Value(userID); uid != nil {
		if i, ok := uid.(int); ok {
			return i
		}
	}
	return 0
}

func GetIsAdmin(ctx context.Context) bool {
	b := ctx.Value(admin)
	if v, ok := b.(bool); ok {
		return v
	}
	return false
}
