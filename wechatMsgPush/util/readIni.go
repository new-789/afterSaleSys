package util

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"gopkg.in/ini.v1"
)

const (
	cfgPath  = "/home/zf007/code/src/afterSaleSys/wechatMsgPush/conf/wxconf.ini"
	CorpId   = "CorpId"   // 企业 di
	AgentId  = "AgentId"  // 应用 ID
	Secret   = "Secret"   // Secret
	MsgType  = "MsgType"  // 消息类型
	PARTNAME = "partName" // 部门名称
)

type BasicInfo struct {
	CorpId   string `json:"corp_id"`  // 企业 di
	AgentId  string `json:"agent_id"` // 应用 ID
	Secret   string `json:"secret"`   // Secret
	PartName string `json:"partname"` // 部门名称
	MsgType  string `json:"msgtype"`  // 消息类型
}

// LoadIni 读取配置文件内容
func LoadIni(section string, keys []string) (basicInfo *BasicInfo, err error) {
	basic := make(map[string]string, 3)
	cfg, err := ini.Load(cfgPath)
	if err != nil {
		fmt.Println("加载配置文件错误...........", err)
		return nil, err
	}
	for _, val := range keys {
		basic[val] = cfg.Section(section).Key(val).String()
	}
	if basicInfo, err = mapToStruct(basic); err != nil {
		fmt.Println("调用 mapToStruct 错误:", err)
		return nil, err
	}
	return
}

// mapToStruct map 转结构体.
func mapToStruct(basic map[string]string) (basicInfo *BasicInfo, err error) {
	basicInfo = &BasicInfo{}
	if err = mapstructure.Decode(basic, basicInfo); err != nil {
		fmt.Println("map 转 struct 发生错误:", err)
		return nil, err
	}
	return
}
