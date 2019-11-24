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

	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponseErrorParams(c)
	}

	userId, _ := strconv.Atoi(c.Param("id"))

	u.ID = uint(userId)

	if err := u.Validate(); err != nil {
		SendResponseErrorValidation(c)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponseErrorDatabase(c)
		return
	}

	SendResponseOK(c)
}
