package src

// 获取 access_token

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

import (
	"encoding/json"
)

const TOKENURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"

type Access struct {
	ErrCode     int    `json:"errcode"`
	ExpiresIn   int    `json:"expires_in"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
}

// GetToken 获取 access_token 方法
func GetToken() (token string, err error) {
	basic, err := GetCfg("token")
	if err != nil {
		fmt.Println("调用 getCfg 方法错误：", err)
		return "nil", err
	}
	reqUrl := fmt.Sprintf(TOKENURL, basic.CorpId, basic.Secret)
	resp, err := http.Get(reqUrl)
	if err != nil {
		fmt.Println("发起获取 token 请求失败：", err)
		return "", err
	}
	defer resp.Body.Close()
	var access = &Access{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取 body 信息失败：", err)
		return "", err
	}
	if err = json.Unmarshal(body, &access); err != nil {
		fmt.Println("json 反序列化错误:", err)
		return "", err
	}
	if access.ErrCode != 0 {
		fmt.Println("json 反序列化错误:", err)
		return "", err
	}
	return access.AccessToken, nil
}
