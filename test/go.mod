module afterSaleSys/test

go 1.16

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.3
	"afterSaleSys/db" v0.0.5
	afterSaleSys/util v0.0.5
    afterSaleSys/wechatMsgPush v1.1.0
    afterSaleSys/wechatMsgPush/src v1.1.0
    afterSaleSys/wechatMsgPush/util v1.1.0
    github.com/gin-contrib/sessions v0.0.3
)

replace (
	"afterSaleSys/db" => "../db"
	afterSaleSys/util => ../util
    afterSaleSys/wechatMsgPush => ../wechatMsgPush
    afterSaleSys/wechatMsgPush/src => ../wechatMsgPush/src
    afterSaleSys/wechatMsgPush/util => ../wechatMsgPush/util
)