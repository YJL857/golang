package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.8.4:6379",
		Password: "111111", // no password set
		DB:       0,        // use default DB
	})

	s := rdb.Get("hello").String()

	fmt.Println(s)

}
