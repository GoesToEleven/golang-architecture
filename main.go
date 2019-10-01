package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancelF := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancelF()
	time.Sleep(50 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("not finished")
	default:
		fmt.Println("finished")
	}
}
