package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/go-redis/redis_rate/v8"
)

// Middleware serving as a limiter.
func Limiter(limiter *redis_rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := limiter.Allow("project:123", redis_rate.PerMinute(3))
		if err != nil {
			panic(err)
		}
		if res.Allowed != true {
			c.Header("X-RateLimit-Retry", res.RetryAfter.String())
			c.AbortWithStatus(http.StatusTooManyRequests)
		}

		c.Header("X-RateLimit-Remaining", strconv.Itoa(res.Remaining))
		c.Header("X-RateLimit-Reset", res.ResetAfter.String())

		fmt.Println(res.Allowed, res.Remaining, res.ResetAfter)

		c.Next()
	}
}

func main() {
	r := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_ = rdb.FlushDB().Err()

	limiter := redis_rate.NewLimiter(rdb)

	// Apply middleware on single route.
	r.GET("/ping", Limiter(limiter), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
