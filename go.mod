module afterSaleSys

go 1.16

require (
	afterSaleSys/db v0.0.5
	afterSaleSys/handler v0.0.5
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.4
	github.com/gorilla/sessions v1.2.1 // indirect
)

replace (
	afterSaleSys/db => ./db
	afterSaleSys/handler => ./handler
	afterSaleSys/util => ./util
	afterSaleSys/wechatMsgPush => ./wechatMsgPush
	afterSaleSys/wechatMsgPush/src => ./wechatMsgPush/src
	afterSaleSys/wechatMsgPush/util => ./wechatMsgPush/util
)
