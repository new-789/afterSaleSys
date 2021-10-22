package src

// 获取部门列表

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	PARTUrl = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s"
)

// PartInfo 部门列表信息
type PartInfo struct {
	Id       int    `json:"id"`
	Parentid int    `json:"parentid"`
	Order    int    `json:"order"`
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
}

type ReturnPartInfo struct {
	ErrCode   int         `json:"errcode"`
	ErrMsg    string      `json:"errmsg"`
	DeartMent []*PartInfo `json:"department"`
}

func GetPartList() (departId int, err error) {
	token, err := GetToken()
	if err != nil {
		fmt.Println("获取部门列表方法中获取 token 失败：", err)
		return 0, err
	}
	urlStr := fmt.Sprintf(PARTUrl, token)
	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Println("获取部门列表请求错误：", err)
		return 0, err
	}
	defer resp.Body.Close()
	var partInfo = &ReturnPartInfo{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取部门列表方法中 ioutil 读取返回信息错误：", err)
		return 0, err
	}
	if err = json.Unmarshal(body, partInfo); err != nil {
		fmt.Println("获取部门列表方法 json 反序列化错误:", err)
		return 0, err
	}
	if partInfo.ErrCode != 0 {
		fmt.Println("获取部门列表失败:", err)
		return 0, err
	}
	part, err := GetCfg("part")
	if err != nil {
		fmt.Println("获取部门列表方法中调用 getCfg 方法错误：", err)
		return 0, err
	}
	for _, val := range partInfo.DeartMent {
		// Name 是否应该放到配置文件？修改起来比较方便
		if val.Name == part.PartName {
			departId = val.Id
			break
		}
	}
	return
}
