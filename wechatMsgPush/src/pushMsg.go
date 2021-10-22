package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 推送消息板块

// PushReturn 推送消息返回数据结构体
type PushReturn struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidPart  string `json:"invalidpart"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

// Content 消息内容结构体
type Content struct {
	Content string `json:"content"`
}

// ReqTextMsg 推送消息请求数据结构体
type ReqTextMsg struct {
	AgentId                int      `json:"agentid"`
	Safe                   int      `json:"safe"`
	EnableIdTrans          int      `json:"enable_id_trans"`
	EnableDuplicateCheck   int      `json:"enable_duplicate_check"`
	DuplicateCheckInterval int      `json:"deplicate_check_interval"`
	ToUser                 string   `json:"touser"` // 可为空
	ToPart                 string   `json:"topart"` // 可为空
	ToTag                  string   `json:"totag"`  // 可为空
	MsgType                string   `json:"msgtype"`
	Markdown               *Content `json:"markdown"`
}

const (
	SAFE                   = 0                  // 是否保密消息
	EnableIdTrans          = 0                  // 是否开启 Id 转译，第三方应用用得上
	EnableDuplicateCheck   = 0                  // 是否开启重复消息检查
	DuplicateCheckInterval = 1800               // 重复消息检查时间间隔
	ContentType            = "application/json" // POST 请求类型
	PUSHURL                = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)

func PushMsgMethod(token string, data *ReqTextMsg) error {
	ret, err := json.Marshal(data)
	if err != nil {
		fmt.Println("推送消息将消息结构体序列化为 []byte 类型失败：", err)
		return err
	}
	reqUrl := fmt.Sprintf(PUSHURL, token)
	resp, err := http.Post(reqUrl, ContentType, bytes.NewReader(ret))
	if err != nil {
		fmt.Println("推送消息发送 post 请求失败：", err)
		return err
	}
	defer resp.Body.Close()
	// 处理返回的数据
	var pushReturn = &PushReturn{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("推送消息读取返回内容失败：", err)
		return err
	}
	if err = json.Unmarshal(body, pushReturn); err != nil {
		fmt.Println("推送消息将 body 反序列化到结构体失败：", err)
		return err
	}
	if pushReturn.ErrCode != 0 {
		fmt.Println("推送消息内容失败....：", err, "======", pushReturn.ErrCode, pushReturn.ErrMsg)
		return err
	}
	fmt.Println("消息推送成功..........................v_v")
	return nil
}
