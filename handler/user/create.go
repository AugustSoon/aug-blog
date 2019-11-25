package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/JumpSama/aug-blog/pkg/errno"
	"github.com/JumpSama/aug-blog/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Infof("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		SendResponseErrorParams(c)
		return
	}

	u := model.UserModel{
		Account:  r.Account,
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
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
