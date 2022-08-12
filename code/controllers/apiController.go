package controllers

import (
	"code/code/model"
	modeladmin "code/code/modelAdmin"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//从数据库获取选项
func GetOption() []string {
	list := []string{}
	model.DB.Raw("SELECT TABLE_NAME FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = 'gin';").Scan(&list)
	return list
}

//获取全部的单词
func GetAllWorld(s string) []model.Word {
	list := []model.Word{}
	model.DB.Table(s).Find(&list)
	return list
}

//从切片里面获取一个随机的元素
func Random(s []model.Word) model.Word {
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())
	length := len(s)
	fmt.Println(length)
	// 产生一个随机元素
	return s[rand.Intn(length)]
}

//强制类型转换
func Transform(list interface{}) []model.Word {
	p, ok := (list).([]model.Word)
	if ok {
		fmt.Println(reflect.TypeOf(p))
	} else {
		fmt.Println("can not convert")
	}
	return p
}

//删除切片元素
func DeleteSlice(a []model.Word, elem string) []model.Word {
	tmp := make([]model.Word, 0, len(a))
	for _, v := range a {
		if v.English != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

//获取所有记录
func GetAllRecord() []modeladmin.Record {
	list := []modeladmin.Record{}
	modeladmin.DB.Raw("SELECT * from record ORDER BY time DESC;").Scan(&list)
	return list
}

//获取排序后的记录
func Order() []modeladmin.Record {
	list := []modeladmin.Record{}
	modeladmin.DB.Raw("SELECT `name`,COUNT(*) time,SUM(`level`) level from record GROUP BY `name`;").Scan(&list)
	return list
}
