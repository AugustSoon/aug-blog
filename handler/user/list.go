package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r model.ListRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		SendResponseErrorParams(c)
		return
	}

	list, count := service.UserList(&r)

	SendResponse(c, nil, ListResponse{
		Total: count,
		List:  list,
	})
}
