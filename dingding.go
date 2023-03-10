package dingding

import (
	"github.com/championlong/dingtalk-sdk/config"
	"github.com/championlong/dingtalk-sdk/model"
)

const (
	dingUrl = "https://oapi.dingtalk.com/robot/send"
)

var dingConfig *config.DingdingConfig

// GetConfig 获取钉钉配置文件
func GetConfig() *config.DingdingConfig {
	return dingConfig
}

// Init 初始化钉钉配置
func Init(config config.DingdingConfig) {
	dingConfig = &config
}

// SendDingMessage 发送钉钉消息
func SendDingMessage(kindRobot string, messageType MsgType, message interface{}, at model.At) error {
	job := &DingMasterJob{}
	job.KindRobot = kindRobot
	job.Url = dingUrl
	job.Msgtype = messageType
	job.Query = dingConfig.DingdingQuery[kindRobot]
	job.At = at
	return job.SendMessage(message)
}
