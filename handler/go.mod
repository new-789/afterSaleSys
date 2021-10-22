module afterSaleSys/handler

go 1.16

require (
	afterSaleSys/db v0.0.5
	afterSaleSys/util v0.0.5
	afterSaleSys/wechatMsgPush v1.1.0
	afterSaleSys/wechatMsgPush/src v1.1.0
	afterSaleSys/wechatMsgPush/util v1.1.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.3
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace (
	afterSaleSys/db => ../db
	afterSaleSys/util => ../util
	afterSaleSys/wechatMsgPush => ../wechatMsgPush
	afterSaleSys/wechatMsgPush/src => ../wechatMsgPush/src
	afterSaleSys/wechatMsgPush/util => ../wechatMsgPush/util
)
