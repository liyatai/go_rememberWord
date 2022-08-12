package main

import (
	"code/code/model"
	"code/code/routers"
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//声明路由
	r := gin.Default()
	//声明静态文件的路由
	r.Static("/static", "./static")
	//声明模板文件的路由
	r.LoadHTMLGlob("templates/*")
	//配置session
	// store := cookie.NewStore([]byte("secret"))
	// r.Use(sessions.Sessions("mysession", store))
	store := cookie.NewStore([]byte("secret"))
	gob.Register([]model.Word{})
	r.Use(sessions.Sessions("mysession", store))
	//实例化自定义路由
	routers.IndexInit(r)
	routers.AdminIndex(r)
	r.Run()
}
