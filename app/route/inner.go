package route

import (
	"github.com/gin-gonic/gin"
	"ucgo/library/isafe"
)

func init() {

}

func InnerRouter(router *gin.Engine) {
	router.Use(InnerGlobalMiddleWare())
	{

	}
}

func InnerGlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 校验来源ip是否为内网ip
		isafe.InnerNetWork()
		// 校验签名
		isafe.InnerSign()
		//前置操作

		c.Next()

		//后置操作
	}
}
