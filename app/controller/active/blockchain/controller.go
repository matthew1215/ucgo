package blockchain

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
	"ucgo/app/model/active/blockchain"
	"ucgo/library/ires"
)

type AddTransferController struct {
	Num string `form:"num" binding:"required"`
}

func AddTransfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params AddTransferController
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			return
		}
		num, _ := strconv.ParseInt(params.Num, 10, 64)
		getDemoInModel := blockchain.AddTransferInModel{num}
		result := getDemoInModel.AddTransfer()
		ires.Result(c, result.Code, result.Msg, result.Data)
		return
	}
}

type IndexController struct{}

func Index(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLFiles("app/controller/active/blockchain/index.html")
		c.HTML(200, "index.html", nil)
	}
}

type PingController struct{}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params PingController
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			return
		}
		var wsupgrader = websocket.Upgrader{
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			HandshakeTimeout: 5 * time.Second,
			// 取消ws跨域校验
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatal("Failed to set websocket upgrade", err)
			return
		}

		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			conn.WriteMessage(t, msg)
		}
		return
	}

}
