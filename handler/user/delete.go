package user

import (
	. "github.com/JumpSama/aug-blog/handler"
	"github.com/JumpSama/aug-blog/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint(userId)); err != nil {
		SendResponseErrorDatabase(c)
		return
	}

	SendResponseOK(c)
}
