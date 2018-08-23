package isafe

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"ucgo/library/icode"
	"ucgo/library/ires"
)

type LoginParams struct {
	Guid     string `form:"guid" binding:"required"`
	Deviceid string `form:"deviceid" binding:"required"`
	Token    string `form:"token" binding:"required"`
	Os       string `form:"os" binding:"required"`
	Time     string `form:"time" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params LoginParams
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.Response(c, icode.NotLogin, icode.StatusText(icode.NotLogin), nil)
			c.Abort()
			return
		}

		// todo 登录校验
	}
}
