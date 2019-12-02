package main

import (
	"fmt"
	"gorm"
	"gorm/generate/mvc"
	"gorm/info"
	"log"
)

// 简单使用ROM映射 返回结果为TABLE或MAP
// 优点是不需要建立结构体、不需要手写SQL
// 缺点是只支持普通的增删改查

func main() {
	// 首先初始化数据库连接 test是数据库名
	mysqlLink := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"
	gorm.Client.Init(mysqlLink)

	// 获取默认数据库信息
	database := info.GetDataBase()

	// 获取默认的数据库查询工具 test是表名
	accountService := mvc.GetService(database, "test")
	// 执行查询
	table, err := accountService.GetList(map[string]string{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(table)
	fmt.Println(table.ToMap())
}
