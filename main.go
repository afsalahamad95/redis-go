package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	hashFields := []string{
		"model", "test",
	}

	// we can set hash values as well
	res, err := client.HSet(ctx, "bike", hashFields).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res2, err := client.HGet(ctx, "bike", "model").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res2)
}
