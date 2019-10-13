package setupFunc

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func RedisMiddleware(client redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", client)
		c.Next()
	}
}
