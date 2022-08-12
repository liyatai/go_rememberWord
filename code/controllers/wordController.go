package controllers

import (
	"code/code/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type WordController struct {
}

//获取表名
func (con WordController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "word.html", gin.H{
		"option": GetOption(),
	})
}

//选择单词本
func (con WordController) Select(c *gin.Context) {
	s := c.Query("key")

	model.SetName(s)
	//获取全部的单词存入切片
	list := GetAllWorld(s)
	//获取一个随机切片的元素
	random := Random(list)
	//获取session
	session := sessions.Default(c)
	//将切片存入session
	session.Set("list", nil)
	session.Set("list", list)
	session.Save()
	c.HTML(200, "random.html", gin.H{
		"option": GetOption(),
		"word":   random,
	})
}
func (con WordController) Next(c *gin.Context) {
	//获取session单词存入切片
	session := sessions.Default(c)
	list := session.Get("list")
	wlist := Transform(list)
	//声明随机单词变量
	var random model.Word
	//获取一个随机切片的元素
	if len(wlist) == 1 {
		random = model.Word{English: "学习完成", Chinese: " "}
	} else {
		random = Random(wlist)
	}

	c.HTML(200, "random.html", gin.H{
		"option": GetOption(),
		"word":   random,
	})
}
func (con WordController) Delete(c *gin.Context) {
	var random model.Word
	//获取请求的参数
	eng := c.Query("eng")
	//获取session单词存入切片
	session := sessions.Default(c)
	list := session.Get("list")
	// 强制类型转换
	wlist := Transform(list)
	//判空
	if len(wlist) == 1 {
		random = model.Word{English: "学习完成", Chinese: " "}
	} else {

		//删除特定元素
		var dlist = []model.Word{}

		for i := 0; i < len(wlist); i++ {
			if wlist[i].English == eng {
				dlist = DeleteSlice(wlist, eng)
			}
		}
		//将切片存入session
		session.Set("list", dlist)
		session.Save()
		//获取一个随机切片的元素
		random = Random(dlist)
	}

	c.HTML(200, "random.html", gin.H{
		"option": GetOption(),
		"word":   random,
	})
}
