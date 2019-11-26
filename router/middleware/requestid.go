package middleware

import (
	"github.com/JumpSama/aug-blog/util"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "X-Request-Id"

		requestId := c.Request.Header.Get(key)

		if requestId == "" {
			requestId = util.GetUUID()
		}

		c.Set(key, requestId)

		c.Writer.Header().Set(key, requestId)

		c.Next()
	}
}
