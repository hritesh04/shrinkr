package database

import (
	"github.com/redis/go-redis/v9"
)

// type Cache struct {
// 	client *redis.Client
// }

var Cache *redis.Client
// func (c *Cache)Get(key string)string{
// 	ctx := context.Background()
// 	val,err := c.client.Get(ctx,key).Result()
// 	if err != nil {
// 		return "Not Cached"
// 	}
// 	return val
// }

func InitCache(){
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}