package session

import "context"

type stringKey string
type intKey int

var userID stringKey
var admin intKey

func SetUserID(ctx context.Context, uID int) context.Context {
	return context.WithValue(ctx, userID, uID)
}

func SetAdminAccess(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

func GetUserID(ctx context.Context) int {
	if v := ctx.Value(userID); v != nil {
		if i, ok := v.(int); ok {
			return i
		}
	}
	return 0
}

func GetAdmin(ctx context.Context) bool {
	if v := ctx.Value(admin); v != nil {
		if i, ok := v.(bool); ok {
			return i
		}
	}
	return false
}
