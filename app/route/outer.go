package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"ucgo/app/controller/active/blockchain"
	"ucgo/app/controller/demo"
	"ucgo/library/isafe"
)

func init() {

}

func OuterRouter(router *gin.Engine) {
	router.Use(OuterGlobalMiddleWare())
	{
		router.POST("/demo", isafe.Post(), isafe.Login(), demo.GetDemo()) // demo
		router.POST("/blockchain/transfer/add", blockchain.AddTransfer()) // 区块链交易
		router.GET("/ping", blockchain.Ping())                            // wb响应
		router.GET("/index", blockchain.Index(router))                    // wb页面
	}
}

func OuterGlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//前置操作
		outerRequestId(c)

		c.Next()

		//后置操作
	}
}

func outerRequestId(c *gin.Context) {
	// todo 有nginx的logid使用nginx的logid
	c.Writer.Header().Set("X-Request-Id", uuid.NewUUID().String())
}
