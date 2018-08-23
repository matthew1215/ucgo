package mongosuser

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	"ucgo/library/imongo"
)

const (
	USER_DB         = "iUser"
	USER_COLLECTION = "user"
)

type GetUserIn struct {
	ID bson.ObjectId
}
type GetUserOut struct {
	ID    bson.ObjectId `bson:"_id"`
	Guid  int64         `bson:"guid"`
	Pguid int64         `bson:"pguid"`
}

// 获取user游标
func (this GetUserIn) GetUser() (*mgo.Session, *mgo.Iter) {
	session := imongo.GetMongosSession()
	session.SetBatch(10)
	session.SetSocketTimeout(60 * time.Second)
	session.SetSyncTimeout(3 * time.Second)
	c := session.DB(USER_DB).C(USER_COLLECTION)
	return session, c.Find(nil).Iter()
}

type GetUserByGuidIn struct {
	Guid string
}
type GetUserByGuidOut struct {
	ID       bson.ObjectId `bson:"_id"`
	Guid     int64         `bson:"guid"`
	Nickname string        `bson:"nickname"`
	Userimg  string        `bson:"userimg"`
	Pguid    int64         `bson:"pguid"`
}

// 通过guid获取用户信息
func (this GetUserByGuidIn) GetUserByGuid() (GetUserByGuidOut, error) {
	session := imongo.GetMongosSession()
	c := session.DB(USER_DB).C(USER_COLLECTION)
	result := GetUserByGuidOut{}
	guid, _ := strconv.ParseInt(this.Guid, 10, 64)
	err := c.Find(bson.M{"guid": guid}).One(&result) //如果查询失败，返回“not found”
	if err != nil {
		return result, err
	}
	return result, err
}
