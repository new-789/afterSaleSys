package handler

import (
	"afterSaleSys/db"
	wechatMsg "afterSaleSys/wechatMsgPush"
	"errors"
	"fmt"
)

// 推送微信消息模块

//var saleNameMap = map[string]string{
//	"陈小兵": "业务", "李珊珊": "业务", "林建立": "业务", "余燕珊": "", "黄静": "美工",
//	"蒋镇霞": "总经办", "吴思燕": "PMC", "康磊": "", "庞辉": "", "郭伟": "", "程途": "商务",
//	"刘燕娟": "商务", "陈文君": "商务", "李广升": "运营", "傅宝锋": "客服", "李冬杏": "运营",
//	"徐锦珊": "商务", "庄雪怡": "商务", "王芷君": "商务",
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
	//default: // 销售为郭伟、庞辉、康磊时提醒谁？
	//	if uname == "郭伟" {
	//		username = "A 链力~" + uname
	//	} else if uname == "康磊" {
	//		username = "链力-" + uname
	//	} else if uname == "庞辉" {
	//		username = uname + "15920648872"
	//	} else if uname == "余燕珊" {
	//		username = "链力-" + uname
	//	}
	//	return ""
	//}

	username = "链力-技术部-朱修建"
	return username
}
