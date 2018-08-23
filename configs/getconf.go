package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configYaml *ConfigYaml = nil

type ConfigYaml struct {
	Inner_port string
	Outer_port string
	Logpath    string
	Mongo      struct {
		Mongos  string
		Replset string
	}
	Redis struct {
		Cluster struct {
			List []string
			Pwd  string
		}
		Cluster2 struct {
			List []string
			Pwd  string
		}
	}
}

func GetConf() *ConfigYaml {
	if configYaml == nil {
		source, err := ioutil.ReadFile("./configs/config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(source, &configYaml)
		if err != nil {
			panic(err)
		}
	}
	return configYaml
}
