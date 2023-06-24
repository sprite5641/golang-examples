package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "is cancelled.")
			return
		default:
			fmt.Println(name, "is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go worker(ctx, "Worker1")
	go worker(ctx, "Worker2")
	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("Main is finished.")
}
