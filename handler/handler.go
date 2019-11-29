package handler

import (
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 响应
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// 正确响应
func SendResponseOK(c *gin.Context) {
	SendResponse(c, nil, nil)
}

// 参数错误响应
func SendResponseErrorParams(c *gin.Context) {
	SendResponse(c, errno.ErrBind, nil)
}

// 数据库错误响应
func SendResponseErrorDatabase(c *gin.Context) {
	SendResponse(c, errno.ErrDatabase, nil)
}

// 参数验证错误响应
func SendResponseErrorValidation(c *gin.Context) {
	SendResponse(c, errno.ErrValidation, nil)
}
