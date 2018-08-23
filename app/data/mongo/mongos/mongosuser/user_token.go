package mongosuser

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"ucgo/library/imongo"
)

const (
	USER_TOKEN_COLLECTION = "user_token"
)

type GetTokenIn struct {
	Guid     string
	Os       string
	Deviceid string
}
type GetTokenOut struct {
	ID    bson.ObjectId `bson:"_id"`
	Token string        `bson:"token"`
}

// 获取用户token
func (this GetTokenIn) GetToken() (GetTokenOut, error) {
	session := imongo.GetMongosSession()
	c := session.DB(USER_DB).C(USER_TOKEN_COLLECTION)
	result := GetTokenOut{}
	guid, _ := strconv.ParseInt(this.Guid, 10, 64)
	err := c.Find(bson.M{"guid": guid, "os": this.Os, "deviceid": this.Deviceid}).One(&result) //如果查询失败，返回“not found”
	if err != nil {
		return result, err
	}
	return result, err
}
