package imongo

import (
	"gopkg.in/mgo.v2"
	"time"
	"ucgo/configs"
)

var MongoReplsetSession *mgo.Session = nil

func GetReplsetSession() *mgo.Session {
	if MongoReplsetSession == nil {
		var err error
		MongoReplsetSession, err = mgo.Dial(configs.GetConf().Mongo.Replset)
		if err != nil {
			panic(err) //直接终止程序运行
		}
		MongoReplsetSession.SetMode(mgo.Monotonic, true)
		//最大连接池默认为4096
		//MongoReplsetSession.SetPoolLimit(2048)	// 根据web服务器数据量来决定连接池的连接数
		MongoReplsetSession.SetSocketTimeout(2 * time.Second)
		MongoReplsetSession.SetSyncTimeout(2 * time.Second)
	}
	return MongoReplsetSession
}
