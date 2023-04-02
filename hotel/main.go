package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	// fake timeout
	//	go func() {
	//		time.Sleep(time.Second * 2)
	//		cancel()
	//	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	case <-time.After(time.Second * 2):
		fmt.Println("Booking completed!")
	}
}
