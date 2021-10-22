package main

import (
	"afterSaleSys/db"
	"afterSaleSys/handler"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := db.InitDb()
	if err != nil {
		return
	}

	handler.R = gin.Default()
	handler.R.LoadHTMLGlob("./view/**/*")
	handler.R.Static("/static/", "./static")

	handler.Store = cookie.NewStore([]byte("LianLi123456"))
	handler.R.Use(sessions.Sessions("mysession", handler.Store))
	//handler.R.Use(handler.MyMiddleware)

	handler.Routers()
	handler.RetRouters()

	if err = handler.R.Run(":6969"); err != nil {
		fmt.Println("服务启动运行失败.......", err)
		return
	}
}
