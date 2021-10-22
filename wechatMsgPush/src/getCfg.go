package src

import (
	"afterSaleSys/wechatMsgPush/util"
	"fmt"
)

// GetCfg 获取配置文件内容
func GetCfg(getType string) (basic *util.BasicInfo, err error) {
	switch getType {
	case "token":
		basic, err = util.LoadIni("wxCont", []string{util.CorpId, util.AgentId, util.Secret, util.MsgType})
		if err != nil {
			fmt.Println("获取配置内容错误：", err)
			return nil, err
		}
	case "part":
		basic, err = util.LoadIni("part", []string{util.PARTNAME})
		if err != nil {
			fmt.Println("获取配置内容错误：", err)
			return nil, err
		}
	}
	return basic, nil
}
