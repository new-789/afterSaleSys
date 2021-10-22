package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//var r *gin.Engine

//func main() {
//	r = gin.Default()
//	store := cookie.NewStore([]byte("LianLi123456"))
//	r.Use(sessions.Sessions("mysession", store))
//
//	r.GET("/hello", test)
//	r.Run(":8000")
//}
//
//func test(c *gin.Context) {
//	se := Init(c)
//	se.SetSe("phone", "world")
//
//	c.JSON(200, gin.H{"hello": se.session.Get("phone")})
//}
//
//type sess struct {
//	session sessions.Session
//}
//
//func Init(c *gin.Context) *sess {
//	store := cookie.NewStore([]byte("LianLi123456"))
//	r.Use(sessions.Sessions("mysession", store))
//	return &sess{
//		session: sessions.Default(c),
//	}
//}
//
//func (s *sess) SetSe(key, val string) bool {
//	s.session.Set(key, val)
//	err := s.session.Save()
//	if err != nil {
//		fmt.Println("se set failed", err)
//		return false
//	}
//	seBool := s.SaveSe()
//	if seBool {
//		return true
//	}
//	return false
//}
//
//func (s *sess) SaveSe() bool {
//	err := s.session.Save()
//	if err != nil {
//		return false
//	}
//	return true
//}

const (
	// REQUESTTYPE 请求类型 POST 请求要用到
	REQUESTTYPE = "application/json"
	// GETTOKENTURL get  请求地址获取 access_token
	GETTOKENTURL     = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	// GETPARTURL 获取部门 Id 请求地址
	GETPARTURL = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%d"
	// GETUSERURL 获取部门成员 Id
	GETUSERURL = ""
	// POSTURL 请求地址，推送消息
	POSTURL     = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"     // post
)

//type WxInterface interface {
//	HttpGetJson() string
//	HttpPostJson()
//}

type WxMethod struct {
	CorpId  string // 企业 di
	AgentId string // 应用 ID
	Secret  string // Secret
}

// ReqRest Get 请求返回的数据结构体
type ReqRest struct {
	ErrCode     int    `json:"errcode"`      // 返回码 0 表示成功,失败返回 40091
	ErrMsg      string `json:"errmsg"`       // 返回码文本描述
	AccessToken string `json:"access_token"` // access_token 凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间
}

// Content 推送消息中 text 的消息数据
type Content struct {
	Content string `json:"content"`
}

// PushMsg 推送消息数据结构体
type PushMsg struct {
	AgentId                int      `json:"agentid"`                  // 企业 id
	Safe                   int      `json:"safe"`                     // 信息是否保密，默认 0 表示可以分享 1 则保密
	EnableIdTrans          int      `json:"enable_id_trans"`          // 是否开启 id 转义，默认 0 第三方应用需要用到
	EnableDuplicateCheck   int      `json:"enable_duplicate_check"`   // 是否开启重复消息检查，默认 0 ，1表示开启
	DuplicateCheckInterval int      `json:"duplicate_check_interval"` // 重复消息的检查时间间隔，默认 1800s ，最大 4 小时
	ToUser                 []string `json:"touser"`                   // 接收的用户 id
	ToParty                []string `json:"toparty"`                  // 接收的部门 Id
	ToTag                  []string `json:"totag"`                    // 接收的标签 Id
	MsgType                string   `json:"msgtype"`                  // 消息类型
	Text                   *Content `json:"text"`                     // 消息内容
}

// PushRspMsg 推送消息返回的数据结构体
type PushRspMsg struct {
	ErrCode      int    `json:"errcode"`       // 返回码
	ErrMsg       string `json:"errmsg"`        // 返回码的描述信息
	InvalidUser  string `json:"invaliduser"`   // 不合法的 userId
	InvalidParty string `json:"invalidparty"`  // 不合法的部门Id
	InvalidTag   string `json:"invalidtag"`    // 不合法的标签 Id
	MsfId        string `json:"msfid"`         // 消息Id，用于撤回应用消息
	ResponseCode string `json:"response_code"` // 消息类型为按钮交互型时，用来更新末班卡片消息的接口
}

func NewWxObj(corpId, agentId, Secret string) *WxMethod {
	return &WxMethod{
		CorpId:  corpId,
		AgentId: agentId,
		Secret:  Secret,
	}
}

// HttpGetAccessToken 发起 get 请求获取access_token 第一步
func (w *WxMethod) HttpGetAccessToken() string {
	reqUrl := fmt.Sprintf(GETTOKENTURL, w.CorpId, w.Secret)
	resp, err := http.Get(reqUrl)
	if err != nil {
		fmt.Println("发起 GET 请求获取 token 信息错误", err)
		return ""
	}
	defer resp.Body.Close()
	temp := &ReqRest{}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	err = json.Unmarshal(all, &temp)
	//err = json.NewDecoder(resp.Body).Decode(&temp)  // 简单写法
	if err != nil {
		return ""
	}
	if temp.ErrCode != 0 {
		fmt.Println("获取 access_token 失败", err)
		return ""
	}
	return temp.AccessToken
}

// HttpGetPartId 获取部门 Id 第二步
func (w *WxMethod) HttpGetPartId() int {
	// TODO 根据返回的 access_token 获取部门 Id
	return 0  // 伪代码,需要返回部门 id
}

// HttpPostJson 发送 POST 推送消息
func (w *WxMethod) HttpPostJson(accessToken string, msg *PushMsg) {
	url := fmt.Sprintf(POSTURL, accessToken)
	res, err := json.Marshal(msg)
	if err != nil {

	}
	http.Post(url, POSTURL, bytes.NewReader(res))
}

func main() {
	//wxObj := NewWxObj("", "", "")
	//accessToken := wxObj.HttpGetAccessToken()
	//msg := &PushMsg{
	//	//AgentId                int      `json:"agentid"`                  // 企业 id
	//	//Safe                   int      `json:"safe"`                     // 信息是否保密，默认 0 表示可以分享 1 则保密
	//	//EnableIdTrans          int      `json:"enable_id_trans"`          // 是否开启 id 转义，默认 0 第三方应用需要用到
	//	//EnableDuplicateCheck   int      `json:"enable_duplicate_check"`   // 是否开启重复消息检查，默认 0 ，1表示开启
	//	//DuplicateCheckInterval int      `json:"duplicate_check_interval"` // 重复消息的检查时间间隔，默认 1800s ，最大 4 小时
	//	//ToUser                 []string `json:"touser"`                   // 接收的用户 id
	//	//ToParty                []string `json:"toparty"`                  // 接收的部门 Id
	//	//ToTag                  []string `json:"totag"`                    // 接收的标签 Id
	//	//MsgType                string   `json:"msgtype"`                  // 消息类型
	//	//Text                   *Content `json:"text"`
	//}
	//wxObj.HttpPostJson(accessToken, msg)
	//db.LdssSltSaleName()
}
