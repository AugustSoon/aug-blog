package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/JumpSama/aug-blog/util"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	account := c.Param("account")

	user, err := model.GetUserByAccount(account)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	} else {
		SendResponse(c, nil, user)
	}
}

// 当前登录用户
func Current(c *gin.Context) {
	id := util.GetCurrentUserId(c)

	user, err := model.GetUserById(id)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	} else {
		SendResponse(c, nil, &model.UserInfo{
			Id:        user.ID,
			Account:   user.Account,
			Username:  user.Username,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
}
