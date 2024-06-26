package base

import (
	"github.com/deng00/go-base/config"
	"log"
)

var configManager = config.Manager{}

const ServiceName = "cs"

func init() {
	// 配置解析
	err := configManager.Init(ServiceName)
	if err != nil {
		panic("init config failed, " + err.Error())
	}
}

func GetConfig() *config.Config {
	c := configManager.GetIns()
	return c
}

func GetArgusConfig() *ArgusConfig {
	var c ArgusConfig
	if err := GetConfig().UnmarshalKey("argus", &c); err != nil {
		log.Fatalf("invalid argus config:%s", err)
	}
	return &c
}

func GetSlackConfig() *SlackConfig {
	var c SlackConfig
	if err := GetConfig().UnmarshalKey("slack", &c); err != nil {
		log.Fatalf("invalid slack config:%s", err)
	}
	return &c
}
