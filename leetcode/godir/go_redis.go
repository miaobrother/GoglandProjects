package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	//连接redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	//客户端进行set操作
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	//客户端实现get操作

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
