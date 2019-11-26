package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/constvar"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/JumpSama/aug-blog/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Infof("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		SendResponseErrorParams(c)
		return
	}

	count := model.GetUserCountByAccount(u.Account, 0)

	if count > 0 {
		SendResponse(c, errno.ErrUserExist, nil)
		return
	}

	if u.Password == "" {
		u.Password = constvar.DefaultPassword
	}

	if err := u.Validate(); err != nil {
		SendResponseErrorValidation(c)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponseErrorDatabase(c)
		return
	}

	rsp := CreateResponse{
		ID:       u.ID,
		Account:  u.Account,
		Username: u.Username,
	}

	SendResponse(c, nil, rsp)
}
