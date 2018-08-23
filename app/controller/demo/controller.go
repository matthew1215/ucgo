package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"ucgo/app/model/demo"
	"ucgo/library/ires"
)

type GetDemoController struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func GetDemo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params GetDemoController
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			return
		}
		getDemoInModel := demo.GetDemoInModel{"test"}
		result := getDemoInModel.GetDemo()
		//ilogger.Info(result.Data)
		ires.Result(c, result.Code, result.Msg, result.Data)
		return
	}
}

type AddDemoControllerParams struct{}

func AddDemoController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params AddDemoControllerParams
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			return
		}
		ires.ResultOk(c, nil)
		return
	}
}

type SetDemoControllerParams struct{}

func SetDemoController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params SetDemoControllerParams
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			return
		}
		ires.ResultOk(c, nil)
		return
	}
}

type DelDemoControllerParams struct{}

func DelDemoController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params DelDemoControllerParams
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			return
		}
		ires.ResultOk(c, nil)
		return
	}
}
