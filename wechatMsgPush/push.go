package wechatMsgPush

import (
	"afterSaleSys/wechatMsgPush/src"
	"fmt"
	"strconv"
)

const (
	SAFE                   = 0       // 是否保密消息
	EnableIdTrans                    // 是否开启 Id 转译，第三方应用用得上
	EnableDuplicateCheck             // 是否开启重复消息检查
	DuplicateCheckInterval = 1800    // 重复消息检查时间间隔
	GETTYPE                = "token" // 消息类型
)

func PushMsgToWechat(username string, message string) error {
	// 获取配置文件内容，传入参数必须是 token，否则出错
	basic, err := src.GetCfg(GETTYPE)
	if err != nil {
		fmt.Println("推送消息获取配置内容失败：", err)
		return err
	}
	// 将企业应该 Id 转换为 int 类型
	agentId, err := strconv.Atoi(basic.AgentId)
	if err != nil {
		fmt.Println("推送消息转换应用id为 int 数据类型失败：", err)
		return err
	}
	// 获取 userId 和 token 传入的参数为想要将消息发送给指定的成员，成员名根据实际场景动态获取
	userId, token, err := src.GetDePartMember(username)
	if err != nil {
		fmt.Println("推送消息调用 GetDePartMember 方法获取 token 和 userId 失败：", err)
		return err
	}
	// 构造推送请求消息数据  TODO 此消息结构体应该从调用处传入，此处仅用来测试
	msg := &src.ReqTextMsg{
		AgentId:                agentId,
		Safe:                   SAFE,
		EnableIdTrans:          EnableIdTrans,
		EnableDuplicateCheck:   EnableDuplicateCheck,
		DuplicateCheckInterval: DuplicateCheckInterval,
		ToUser:                 userId,
		ToPart:                 "",
		ToTag:                  "",
		MsgType:                basic.MsgType,
		Markdown: &src.Content{
			Content: message,
		},
	}
	if err = src.PushMsgMethod(token, msg); err != nil {
		fmt.Println("消息推送失败")
		return err
	}
	//fmt.Println("消息推送成功")
	return nil
}

// 调用方式示例
//func main() {
//	/*
//		调用步骤：
//			1. 去配置文件修改相关信息
//			2. 将上面的 PushMsgToWechat 方法拷贝到一个固定模块中
//			3. 初始化消息内容 src.Content{Content: "需要发送的消息内容"}
//			4. 根据实际情况获取需要推送的成员名称
//			5. 调用 PushMsgToWechat 方法并传参，参数一为接受消息的成员名，参数二为 *src.Content 结构体即需要发送的消息内容
//	*/
//	// 定义消息数据
//	msg := "测试内容来了,请让道....\n<a href='https://www.bing.com/'>这是链接</a>\n点击即可前往,恭喜您成功了，几天的功夫没有浪费！"
//	if err := PushMsgToWechat("悠悠破解", msg); err != nil {
//
//	}
//}
