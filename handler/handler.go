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

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendResponseOK(c *gin.Context) {
	SendResponse(c, nil, nil)
}

func SendResponseErrorParams(c *gin.Context) {
	SendResponse(c, errno.ErrBind, nil)
}

func SendResponseErrorDatabase(c *gin.Context) {
	SendResponse(c, errno.ErrDatabase, nil)
}

func SendResponseErrorValidation(c *gin.Context) {
	SendResponse(c, errno.ErrValidation, nil)
}
