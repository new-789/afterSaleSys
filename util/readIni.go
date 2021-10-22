package util

import (
	"gopkg.in/ini.v1"
)

// ReadIni 读取 ini 配置文件内容
func ReadIni(section string, keys []string) (map[string]string, error) {
	cfgMap := make(map[string]string, 5)
	cfg, err := ini.Load("/home/zf007/code/src/afterSaleSys/cfg/config.ini")
	if err != nil {
		return nil, err
	}
	for _, value := range keys {
		cfgMap[value] = cfg.Section(section).Key(value).String()
	}
	return cfgMap, nil
}
