package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/JumpSama/aug-blog/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Infof("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var u model.User

	if err := c.Bind(&u); err != nil {
		SendResponseErrorParams(c)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	userId := uint(id)

	count := model.GetUserCountByAccount(u.Account, userId)

	if count > 0 {
		SendResponse(c, errno.ErrUserExist, nil)
		return
	}

	if err := u.Validate(); err != nil {
		SendResponseErrorValidation(c)
		return
	}

	if u.Password != "" {
		if err := u.Encrypt(); err != nil {
			SendResponse(c, errno.ErrEncrypt, nil)
			return
		}
	}

	var user model.User
	user.ID = userId

	if err := user.Update(&u); err != nil {
		SendResponseErrorDatabase(c)
		return
	}

	SendResponseOK(c)
}
