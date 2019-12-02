package main

import (
	"fmt"
	"gorm"
	"gorm/generate/mvc"
	"gorm/info"
)

// 生成SQL信息，可写出到文件

func main() {
	// 首先初始化数据库连接 test是数据库名
	mysqlLink := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"
	gorm.Client.Init(mysqlLink)

	// 获取默认数据库信息
	database := info.GetDataBase()

	// 生成SQL XML
	xml := mvc.GetMapperXML(database, "t_test")
	fmt.Println(xml)
}
