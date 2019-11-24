package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context)  {
	account := c.Param("account")

	user, err := model.GetUserByAccount(account)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}

	SendResponse(c, nil, user)
}