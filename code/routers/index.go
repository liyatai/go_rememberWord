package routers

import (
	"code/code/controllers"

	"github.com/gin-gonic/gin"
)

func IndexInit(r *gin.Engine) {

	router := r.Group("/word")
	router.GET("/test", func(ctx *gin.Context) {
		list := controllers.GetOption()
		ctx.JSON(200, list)
	})
	//首页渲染
	router.GET("/index", controllers.WordController{}.Index)
	//切换渲染
	router.GET("/next", controllers.WordController{}.Next)
	//删除单词
	router.GET("/delete", controllers.WordController{}.Delete)
	router.GET("/pronunciation", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	router.GET("/admin", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	router.GET("/select", controllers.WordController{}.Select)
}
