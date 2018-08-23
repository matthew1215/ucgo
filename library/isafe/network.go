package isafe

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InnerNetWork() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 校验内网ip
		ip := strings.Split(c.ClientIP(), ".")
		innerIpSeg := map[string]int{
			"10":  1,
			"127": 1,
			"172": 1,
			"192": 1,
		}
		if _, ok := innerIpSeg[ip[0]]; !ok {
			c.String(http.StatusNotFound, "404 page not found")
			c.Abort()
			return
		}
	}

}
