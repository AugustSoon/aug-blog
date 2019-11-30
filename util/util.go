package util

import (
	"github.com/JumpSama/aug-blog/pkg/constvar"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	id := uuid.NewV4()
	return id.String()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")

	if !ok {
		return ""
	}

	if requestId, ok := v.(string); ok {
		return requestId
	}

	return ""
}

func GetCurrentUserId(c *gin.Context) uint {
	v := c.GetInt(constvar.CurrentUserIdKey)

	return uint(v)
}
