package iutil

import (
	"github.com/gin-gonic/gin"
)

var guid string
var deviceid string
var proid string
var path string

func ParamsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		guid = c.DefaultQuery("guid", "")
		deviceid = c.DefaultQuery("deviceid", "")
		proid = c.DefaultQuery("proid", "")
		path = c.Request.URL.Path
	}
}

func GetGuid() string {
	return guid
}

func GetDeviceid() string {
	return deviceid
}

func GetProid() string {
	return proid
}

func GetPath() string {
	return path
}
