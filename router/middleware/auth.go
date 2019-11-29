package middleware

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/pkg/constvar"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/JumpSama/aug-blog/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := token.ParseRequest(c)

		if err != nil {
			SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set(constvar.LoginUserIdKey, u.ID)

		c.Next()
	}
}
