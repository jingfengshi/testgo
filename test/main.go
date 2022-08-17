package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	ctx := context.WithValue(context.Background(), "key", "value01")
	go func() {
		for {
			_ = context.WithValue(ctx, "key", "value02")
		}
	}()
	go func() {
		for {
			_ = context.WithValue(ctx, "key", "value03")
		}
	}()
	go func() {
		for {
			fmt.Println(ctx.Value("key"))
		}
	}()
	go func() {
		for {
			fmt.Println(ctx.Value("key"))
		}
	}()
	time.Sleep(3 * time.Second)
}
