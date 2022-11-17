package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// client
	c := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// ctx
	ctx := context.Background()

	// set
	err := c.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		panic(err)
	}

	// get
	v, err := c.Get(ctx, "hello").Result()
	if err == redis.Nil {
		fmt.Println("key not exist")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
