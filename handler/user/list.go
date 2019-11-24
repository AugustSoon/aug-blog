package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r model.ListRequest

	if err := c.Bind(&r); err != nil {
		SendResponseErrorParams(c)
		return
	}

	list, count := model.GetUserList(&r)

	SendResponse(c, nil, ListResponse{
		Total: count,
		List:  list,
	})
}
