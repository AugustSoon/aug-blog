package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	id := uuid.NewV4()
	return fmt.Sprintf("%s", id)
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
