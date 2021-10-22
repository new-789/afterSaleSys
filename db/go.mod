module afterSaleSys/db

go 1.16

require (
	afterSaleSys/util v0.0.5
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.3.4
	afterSaleSys/wechatMsgPush v1.1.0
	afterSaleSys/wechatMsgPush/src v1.1.0
	afterSaleSys/wechatMsgPush/util v1.1.0
)

replace (
	afterSaleSys/util => ../util
	afterSaleSys/wechatMsgPush => ../wechatMsgPush
	afterSaleSys/wechatMsgPush/src => ../wechatMsgPush/src
	afterSaleSys/wechatMsgPush/util => ../wechatMsgPush/util
)