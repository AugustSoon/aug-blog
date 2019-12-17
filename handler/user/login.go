package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/auth"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/JumpSama/aug-blog/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserLogin

	if err := c.ShouldBindJSON(&u); err != nil {
		SendResponseErrorParams(c)
		return
	}

	user, err := model.GetUserByAccount(u.Account)

	if err != nil {
		SendResponse(c, errno.ErrLogin, nil)
		return
	}

	if !auth.Compare(user.Password, u.Password) {
		SendResponse(c, errno.ErrLogin, nil)
		return
	}

	t, e := token.Sign(token.Context{
		ID:      user.ID,
		Account: user.Account,
	})

	if e != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
