package route

import (
	"fmt"
	"os"
	"reflect"
	"ucgo/app/script/blockchain/watchtransferscript"
)

// go run ucgo.go script watchtransfer
// ./ucgo.go script watchtransfer
var ScriptRouteMap = map[string]interface{}{
	"watchtransfer": watchtransferscript.Do, // 区块链监控交易记录
}

// 执行脚本
func Script() bool {
	scriptArg := os.Args[1:]
	if len(scriptArg) >= 1 && scriptArg[0] == "script" {
		if len(scriptArg) >= 2 {
			if _, ok := ScriptRouteMap[scriptArg[1]]; ok {
				script := scriptArg[1]
				scriptCall(ScriptRouteMap, script)
				return true
			} else {
				fmt.Println("脚本不存在")
				return true
			}
		} else {
			fmt.Println("脚本名不能为空")
			return true
		}
	}
	return false
}

func scriptCall(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	if in != nil {
		result = f.Call(in)
	}
	return
}
