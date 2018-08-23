package isafe

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
	"ucgo/library/ires"
	"ucgo/library/iutil"
)

var innerSignConf = map[string]map[string]string{
	"1": {
		"partner": "H5",
		"salt":    "123",
	},
}

type InnerSignParams struct {
	Keyid string `form:"keyid" binding:"required"`
	Time  string `form:"time" binding:"required"`
	Sign  string `form:"sign" binding:"required"`
}

// 校验签名
func InnerSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证参数是否缺失
		var params InnerSignParams
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			c.Abort()
			return
		}
		if _, ok := innerSignConf[params.Keyid]; !ok {
			ires.ResultParamsErr(c)
			c.Abort()
			return
		}
		salt := innerSignConf[params.Keyid]["salt"]
		signatureString := "keyid=" + params.Keyid + "&time=" + params.Time
		genSign := iutil.Md5Encrypt(strings.Join([]string{signatureString, salt}, ""))
		if genSign != params.Sign {
			ires.ResultParamsErr(c)
			c.Abort()
			return
		}
	}
}
