package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 获取部门成员(即企业微信中的员工)

// Member 成员信息
type Member struct {
	UserId     string `json:"userid"`
	Name       string `json:"name"`
	OpenUserId string `json:"open_userid"`
	DepartMent []int  `json:"department"`
}

// PartMemberInfo 成员返回信息
type PartMemberInfo struct {
	ErrCode  int       `json:"errcode"`
	ErrMsg   string    `json:"errmsg"`
	UserList []*Member `json:"userlist"`
}

const (
	REQMEMURL  = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d&fetch_child=%d"
	FETCHCHILD = 0
)

func GetDePartMember(userName string) (userId, token string, err error) {
	token, err = GetToken()
	if err != nil {
		fmt.Println("获取部门成员调用 getToken 方法获取 token 错误:", err)
		return "","", err
	}
	partId, err := GetPartList()
	if err != nil {
		fmt.Println("获取部门成员调用 getPartList 方法获取 partId 错误:", err)
		return "","", err
	}
	reqUrl := fmt.Sprintf(REQMEMURL, token, partId, FETCHCHILD)
	resp, err := http.Get(reqUrl)
	if err != nil {
		fmt.Println("获取部门成员发起 get 请求失败：", err)
		return "","", err
	}
	defer resp.Body.Close()
	var partMemInfo = &PartMemberInfo{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取部门成员中 ioutil 读取 body 信息失败:", err)
		return "","", err
	}
	if err = json.Unmarshal(body, &partMemInfo); err != nil {
		fmt.Println("获取部门成员 json 序列化数据到结构体失败:", err)
		return "","", err
	}
	for _, val := range partMemInfo.UserList {
		if val.Name == userName {
			userId = val.UserId
			break
		}
	}
	return
}
