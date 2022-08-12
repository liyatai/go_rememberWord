package routers

import (
	"code/code/controllers"

	"github.com/gin-gonic/gin"
)

func AdminIndex(r *gin.Engine) {
	router := r.Group("/admin")
	router.GET("/index", controllers.AdminController{}.Aindex)
	router.GET("/add", controllers.AdminController{}.AdminAdd)
	router.GET("/order", controllers.AdminController{}.AdminOrder)
}
