package ires

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"ucgo/library/icode"
)

type ResponseStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const JsonpCallback = "_callback"

// 返回(尽量不直接调用)
func Response(c *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	response := ResponseStruct{Code: code, Msg: msg, Data: data}
	jsonpCallback, _ := c.GetQuery(JsonpCallback)
	if jsonpCallback == "" {
		c.JSON(http.StatusOK, response)
	} else {
		c.Header("content-type", "application/json;charset=utf-8")
		b, _ := json.Marshal(response)
		c.String(http.StatusOK, jsonpCallback+"("+string(b)+")")
	}
	return
}

// ---------------------------以下为controller使用-------------------------------------
// 返回
func Result(c *gin.Context, code int, msg string, data interface{}) {
	Response(c, code, msg, data)
	return
}

// 返回成功
func ResultOk(c *gin.Context, data interface{}) {
	Response(c, icode.Ok, icode.StatusText(icode.Ok), data)
	return
}

// 返回警告
func ResultWarning(c *gin.Context, code int, msg string, data interface{}) {
	Response(c, code, msg, data)
	return
}

// 返回参数错误
func ResultParamsErr(c *gin.Context) {
	Response(c, icode.ParamsError, icode.StatusText(icode.ParamsError), nil)
	return
}

// ---------------------------以下为model使用-------------------------------------------
// 返回成功
func Ok(data interface{}) ResponseStruct {
	return ResponseStruct{icode.Ok, icode.StatusText(icode.Ok), data}
}

// 返回警告
func Warning(code int, msg string, data interface{}) ResponseStruct {
	return ResponseStruct{code, msg, data}
}
