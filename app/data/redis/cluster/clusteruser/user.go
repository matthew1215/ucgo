package clusteruser

import (
	"fmt"
	"strings"
	"time"
	"ucgo/library/iredis"
	"ucgo/library/iutil"
)

const REDIS_USER_STRING_USER_TOKEN = "user:user_token:%s" //用户token

type GetTokenIn struct {
	Guid     string
	Os       string
	Deviceid string
}

// 获取token
func (this GetTokenIn) GetToken() string {
	key := fmt.Sprintf(REDIS_USER_STRING_USER_TOKEN, iutil.Md5Encrypt(strings.Join([]string{this.Guid, this.Os, this.Deviceid}, "")))
	return iredis.GetClusterSession().Get(key).Val()
}

type SetTokenIn struct {
	Guid     string
	Os       string
	Deviceid string
	Token    string
}

// 设置用户token
func (this SetTokenIn) SetToken() {
	key := fmt.Sprintf(REDIS_USER_STRING_USER_TOKEN, iutil.Md5Encrypt(strings.Join([]string{this.Guid, this.Os, this.Deviceid}, "")))
	iredis.GetClusterSession().Set(key, this.Token, time.Minute)
}
