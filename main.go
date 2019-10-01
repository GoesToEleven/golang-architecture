package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancelF := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancelF()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("work didn't finish")
	default:
		fmt.Println("work finished")
	}

}
