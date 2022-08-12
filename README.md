# go_rememberWord
基于go语言开发的记单词的web小项目，web框架用的gin，sql框架用的gorm，数据库用的mariaDB，兼容MySQL
# 初始化
由于我个人使用的云服务器，在这个仓库中的数据库连接的代码的url都是省略了的，所以数据库连接的参数要自行设置(服务器太小了撑不住大家造，而且主要怕有大哥ddos)
需要修改的文件有两个
**1**  /model/sql.go
**2**  /modelAdmin/sql.go
## 修改model下的sql.go
密码、ip地址啥的就自己填写，这里得数据库名称也随便填写,这个数据库用来存放单词得数据
```go
package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:密码@!@tcp(ip地址)/数据库名称?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

```
## 修改modelAdmin下的sql.go
密码、ip地址啥的就自己填写，这里得数据库名称也随便填写，这个数据库用来记录学习记录，注意这个数据库下必须建立一个record名字的表
```go
package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:密码@!@tcp(ip地址)/数据库名称?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
```
# 建表格式
## 单词表
名字随便取，只需要两列，English和Chinese，然后都是vchar，建议 30个左右一张表，给表取名最好有数字索引，首先多了一次性记不完，其次数据添加多了可能会崩。
![image](https://user-images.githubusercontent.com/102398022/184393820-1132dc17-7a04-48d2-b953-aa2f2dfd09e5.png)
## 记录表
记录表只能建立一个record名字命名的表


