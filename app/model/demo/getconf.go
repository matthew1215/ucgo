package demo

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configYaml *ConfigYaml = nil

type ConfigYaml struct {
}

func GetConf() *ConfigYaml {
	if configYaml == nil {
		source, err := ioutil.ReadFile("./app/model/demo/config.yaml")
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
