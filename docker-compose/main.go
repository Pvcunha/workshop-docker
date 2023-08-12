package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func getHitCount() int64 {

	var retries int = 5

	for {
		count, err := rdb.Incr(ctx, "hits").Result()
		if err == nil {
			return count
		}

		if err == redis.Nil {
			return 1
		}

		if err != nil {
			if retries == 0 {
				panic(err)
			}
			retries--
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		count := getHitCount()
		c.String(200, fmt.Sprintf("Hello, docker! I was hitted %d times\n", count))
	})

	r.Run(":8080")
}
