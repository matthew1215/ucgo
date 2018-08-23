package iredis

import (
	"github.com/go-redis/redis"
	"ucgo/configs"
)

var ClusterSession *redis.ClusterClient = nil
var Cluster2Session *redis.ClusterClient = nil

func GetClusterSession() *redis.ClusterClient {
	if ClusterSession == nil {
		ClusterSession = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    configs.GetConf().Redis.Cluster.List,
			Password: configs.GetConf().Redis.Cluster.Pwd,
		})
		//pong, err := ClusterSession.Ping().Result()
		//fmt.Println(reflect.TypeOf(ClusterSession), pong, err)
	}
	return ClusterSession
}

func GetCluster2Session() *redis.ClusterClient {
	if Cluster2Session == nil {
		Cluster2Session = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    configs.GetConf().Redis.Cluster2.List,
			Password: configs.GetConf().Redis.Cluster2.Pwd,
		})
		//pong, err := Cluster2Session.Ping().Result()
		//fmt.Println(reflect.TypeOf(Cluster2Session), pong, err)
	}
	return Cluster2Session
}
