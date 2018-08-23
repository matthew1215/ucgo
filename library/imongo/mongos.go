package imongo

import (
	"gopkg.in/mgo.v2"
	"time"
	"ucgo/configs"
)

var MongoMongosSession *mgo.Session = nil

func GetMongosSession() *mgo.Session {
	if MongoMongosSession == nil {
		var err error
		MongoMongosSession, err = mgo.Dial(configs.GetConf().Mongo.Mongos)
		if err != nil {
			panic(err) //直接终止程序运行
		}
		MongoMongosSession.SetMode(mgo.Monotonic, true)
		//最大连接池默认为4096
		//MongoMongosSession.SetPoolLimit(2048)	// 根据web服务器数据量来决定连接池的连接数
		MongoMongosSession.SetSocketTimeout(2 * time.Second)
		MongoMongosSession.SetSyncTimeout(2 * time.Second)
	}
	return MongoMongosSession
}
