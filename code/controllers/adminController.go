package controllers

import (
	modeladmin "code/code/modelAdmin"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

//初始化后台
func (con AdminController) Aindex(c *gin.Context) {
	list := GetAllRecord()
	c.HTML(200, "admin.html", gin.H{
		"list":   list,
		"option": GetOption(),
	})
}

//后台添加数据
func (con AdminController) AdminAdd(c *gin.Context) {
	name := c.Query("name")
	level := c.Query("level")
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//生成结构体
	record := modeladmin.Record{Name: name, Level: level, Time: timeStr}
	//存入数据库
	modeladmin.DB.Create(&record)
	//获取全部数据
	list := GetAllRecord()
	c.HTML(200, "admin.html", gin.H{
		"list":   list,
		"option": GetOption(),
	})

}

//后台数据排序
func (con AdminController) AdminOrder(c *gin.Context) {
	list := Order()
	c.HTML(200, "admin.html", gin.H{
		"list":   list,
		"option": GetOption(),
	})

}
