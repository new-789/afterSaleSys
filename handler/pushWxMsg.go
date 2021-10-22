package handler

import (
	"afterSaleSys/db"
	wechatMsg "afterSaleSys/wechatMsgPush"
	"errors"
	"fmt"
)

// 推送微信消息模块

//var saleNameMap = map[string]string{
//	该 map 用来定义需要推送的,企业微信内部的人名(个人姓名),可以是多个
//}

// PushMsgToWechat 开始推送消息数据
func pushMsgToWechat(ldssNum, mate string, msg *string) (err error) {
	var username string
	// 1. 根据快递单号查询销售人员 Id，再根据根据销售人员Id 查询销售人员姓名
	username, err = db.LdssSltSaleName(ldssNum,mate)
	if err != nil {
		fmt.Println("=========== 发送微信通知时，根据快递单号查询用户名错误 ==========", err)
		return err
	}
	username = joinName(username)
	fmt.Println("=======================================>", username)
	if username == "" {
		return errors.New("数据类型错误，不支持的用户名")
	}
	// 5. 调用 wechatMsgPush 模块中的 `PushMsgToWechat` 方法开始推送消息
	err = wechatMsg.PushMsgToWechat(username, *msg)
	if err != nil {
		fmt.Println("调用 wechatMsgPush 模块中的 PushMsgToWechat 方法推送消息失败：", err)
		return err
	}
	return nil
}

// 根据销售人员名称拼接企业微信中的成员格式
func joinName(uname string) (username string) {
	// 3. 按照销售人员和业务人员拼成员名称
	// 	3.1: 新建两个 slice 数据结构存储商务和业务人员

	//switch saleNameMap[uname] {
	//case "商务": // 	3.2：若查询到的人员在商务 MAP 中则拼接为 "链力-商务-人名"
	//	if uname == "程途" {
	//		username = "链力-商务-程小姐"
	//	} else {
	//		username = "链力-商务-" + uname
	//	}
	//case "业务": // 	3.3：若查询到的人员在业务 MAP 中则拼接为 "链力-业务-人名"
	//	username = "链力-业务-" + uname
	//case "总经办":
	//	username = "A链力-总经办-蒋小姐"
	//case "美工":
	//	username = "链力-美工-" + uname
	//case "PMC":
	//	username = "链力-PMC-" + uname
	//case "客服":
	//	username = "链力-客服-" + uname
	//case "运营":
	//	username = "链力-运营-" + uname
	//default: // 当销售人员在企业微信中的名称比较特殊时可在该节点下单独进行拼接？
	//	return ""
	//}

	username = "测试用名称"
	return username
}
