package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client
func main(){
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Success")
	}
}
