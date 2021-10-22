module afterSaleSys/util

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.62.0
	afterSaleSys/wechatMsgPush v1.1.0
	afterSaleSys/wechatMsgPush/src v1.1.0
	afterSaleSys/wechatMsgPush/util v1.1.0
)

replace (
	afterSaleSys/wechatMsgPush => ../wechatMsgPush
	afterSaleSys/wechatMsgPush/src => ../wechatMsgPush/src
	afterSaleSys/wechatMsgPush/util => ../wechatMsgPush/util
)