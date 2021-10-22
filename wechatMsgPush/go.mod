module afterSaleSys/wechatMsgPush

go 1.17

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	afterSaleSys/wechatMsgPush/src v1.1.0
	afterSaleSys/wechatMsgPush/util v1.1.0 // indirect
)

replace (
	afterSaleSys/wechatMsgPush/src => ./src
	afterSaleSys/wechatMsgPush/util => ./util
)
