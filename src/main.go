package main

import (
	"net/http"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

var rdb *redis.Client

var key string = "echo"
var field string = "v1"

func init() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:       addr,
		DB:         0,
		PoolSize:   3,
		MaxRetries: 3,
	})

}

func main() {
	e := echo.New()
	e.GET("/incr", func(c echo.Context) error {
		res, err := rdb.HGet(key, field).Result()
		if err == redis.Nil {
			res = ""
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, H{"msg": err})
		}

		err = rdb.HIncrBy(key, field, 1).Err()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{"msg": err})
		}
		if res == "" {
			return c.JSON(http.StatusOK, H{"msg": "hello world", "version": field})
		}
		return c.JSON(http.StatusOK, H{"msg": "hello world!", "count": res, "version": field})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
