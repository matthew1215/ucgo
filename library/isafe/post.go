package isafe

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"ucgo/library/ires"
	"ucgo/library/iutil"
)

type PostParams struct {
	EncryptKey   string `form:"encrypt_key" binding:"required"`
	SecureParams string `form:"secure_params" binding:"required"`
	Os           string `form:"os" binding:"required"`
	Gv           string `form:"gv" binding:"required"`
	Proid        string `form:"proid" binding:"required"`
	Deviceid     string `form:"deviceid" binding:"required"`
	Guid         string `form:"guid"`
	Time         string `form:"time" binding:"required"`
}

// 请求过期时间
const POST_EXPIRE_TIME int = 86400

// * postConf配置的版本注意有由高到低 否则解密错误
// 私钥固定格式
const POST_RSA_PRIVATE_KEY = "./library/isafe/keystore/%s/private_%s.key"

// * postConf配置的版本注意有由高到低 否则解密错误
var postConf = map[string]map[string][]string{
	iutil.PROID_NORMAL: {
		iutil.DEVICE_IOS: {
			"2.0.0|v2",
			"1.0.0|v1",
		},
		iutil.DEVICE_ANDROID: {
			"2.0.0|v2",
			"1.0.0|v1",
		},
	},
}

// * postConf配置的版本注意有由高到低 否则解密错误
var postVersionConf = map[string]map[string]string{
	"v2": {
		"salt": "123",
	},
	"v1": {
		"salt": "123",
	},
}

// * postConf配置的版本注意有由高到低 否则解密错误
func Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证参数是否缺失
		var params PostParams
		if c.ShouldBindWith(&params, binding.Form) != nil {
			ires.ResultParamsErr(c)
			c.Abort()
			return
		}
		version := getVersion(params.Proid, iutil.GetOS(params.Os), params.Gv)
		if version == "" {
			ires.ResultParamsErr(c)
			c.Abort()
			return
		}
		// 私钥
		privateFilePath := fmt.Sprintf(POST_RSA_PRIVATE_KEY, version, iutil.GetOS(params.Os))
		// 盐值
		salt := postVersionConf[version]["salt"]
		// 校验post加密参数
		doPost(c, params, privateFilePath, salt)
	}
}

// 获取版本
func getVersion(proid string, os string, gv string) string {
	versionMap := postConf[proid][os]
	for _, versionSli := range versionMap {
		versionStr := strings.Split(versionSli, "|")
		if iutil.VersionCompare(gv, versionStr[0], ">=") {
			return versionStr[1]
		}
	}
	return ""
}

// 校验post加密参数
func doPost(c *gin.Context, params PostParams, privateFilePath string, salt string) {
	// base解码参数encrypt_key
	decodeEncryptKeyByte, err := base64.StdEncoding.DecodeString(params.EncryptKey)
	if err != nil {
		log.Fatalln(err)
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
	// base解码参数secure_params
	decodeSecureParamsByte, err := base64.StdEncoding.DecodeString(params.SecureParams)
	if err != nil {
		log.Fatalln(err)
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
	// 获取rsa私钥
	privateKey, err := ioutil.ReadFile(privateFilePath)
	if err != nil {
		log.Fatalln(err)
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
	// rsa解密
	desKey, err := iutil.RsaDecrypt(privateKey, decodeEncryptKeyByte)
	if err != nil {
		log.Fatalln(err)
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
	// des解密
	desByte, err := iutil.DesDecrypt(decodeSecureParamsByte, desKey, desKey)
	if err != nil {
		log.Fatalln(err)
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
	// 验证签名
	verifyStatus, dataMap := verifySign(c, desByte, salt)
	if verifyStatus == false {
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
	// 参数校验
	if verifyParams(params, dataMap) == false {
		ires.ResultParamsErr(c)
		c.Abort()
		return
	}
}

// 参数校验
func verifyParams(params PostParams, dataMap map[string]string) bool {
	h5Time, _ := strconv.Atoi(dataMap["h5_time"])
	clientTime, _ := strconv.Atoi(dataMap["client_time"])
	timestamp := iutil.GetTimeStamp()
	// h5时间不能超过一天
	if dataMap["h5_time"] == "" || h5Time-timestamp > POST_EXPIRE_TIME {
		return false
	}
	// 客户端时间不能超过一天
	if dataMap["client_time"] == "" || clientTime-timestamp > POST_EXPIRE_TIME {
		return false
	}
	// 已登陆用户 标参guid与加密guid必须相等
	if (dataMap["guid"] != "" && params.Guid != dataMap["guid"]) ||
		(params.Guid != "" && params.Guid != dataMap["guid"]) {
		return false
	}
	// 标参deviceid与加密deviceid必须相等
	if dataMap["deviceid"] == "" || params.Deviceid != dataMap["deviceid"] {
		return false
	}
	// 标参的时间跟h5时间与客户端时间都不相同
	if (dataMap["time"] != "h5_time") && (dataMap["time"] != "client_time") {
		return false
	}
	return true
}

// 验证签名
func verifySign(c *gin.Context, desByte []byte, salt string) (bool, map[string]string) {
	// 根据&分隔
	desArr := strings.Split(string(desByte), "&")
	// 原签名
	var sign string
	// 排序key
	keysSlice := make([]string, len(desArr))
	i := 0
	dataMap := make(map[string]string)
	for _, v := range desArr {
		value := strings.Split(string(v), "=")
		// 参数复制
		c.Request.Form.Set(value[0], value[1])
		if value[0] == "sign" {
			sign = value[1]
			keysSlice = append(keysSlice[:i], keysSlice[i+1:]...)
			continue
		}
		keysSlice[i] = value[0]
		dataMap[value[0]] = value[1]
		i++
	}
	// 按照key排序
	sort.Strings(keysSlice)
	dataSlice := make([]string, len(keysSlice))
	for k, v := range keysSlice {
		dataSlice[k] = strings.Join([]string{v, dataMap[v]}, "=")
	}
	// 不含签名的原数据串
	noSignStr := strings.Join(dataSlice, "&")
	// 生成的签名
	genSign := iutil.Md5Encrypt(strings.Join([]string{noSignStr, salt}, ""))
	if genSign != sign {
		return false, make(map[string]string)
	}
	return true, dataMap
}
