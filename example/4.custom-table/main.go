package main

import (
	"fmt"
	"gorm"
)

// 使用自定义SQL的查询 返回结果为TABLE或MAP
// 优点是不需要建立结构体

func main() {
	// 这里的SQL可以使用变量，也可以从文件中读取
	sql := `
		<xml>
    		<sql>
        		<key>GetList</key>
        		<value>
            		SELECT
                		id, title
            		FROM
                		t_test
        		</value>
    		</sql>
			<sql>
        		<key>GetCount</key>
        		<value>
            		SELECT
                		COUNT(1)
            		FROM
                		t_test
        		</value>
    		</sql>
			<sql>
        		<key>Delete</key>
        		<value>
            		Delete
            		FROM
                		t_test
					WHERE id = 4
        		</value>
    		</sql>
		</xml>
	`
	// 首先初始化数据库连接 test是数据库名
	mysqlLink := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"
	gorm.Client.Init(mysqlLink)

	// 解析SQL
	xmlConfig := gorm.XmlConfig(sql)
	// 创建管理者
	manager := gorm.GetTableManager(xmlConfig)
	// 当有多个数据源时可以使用 切换数据源
	manager.Engine.Use(gorm.DefaultClientAlias)
	// manager有三个方法
	// 查询数据集 参数SQL的别名 查询条件
	table, err := manager.Query("GetList", map[string]string{})
	if err == nil {
		fmt.Println(table)
	}
	// 查询Count
	number, err := manager.Count("GetCount", map[string]string{})
	if err == nil {
		fmt.Println(number)
	}
	// 执行语句
	number1, number2, err := manager.Exec("Delete", "")
	if err == nil {
		fmt.Println(number1, number2)
	}
	// number1 最后一次插入ID 当执行添加语句时有效
	// number2 影响的行数 当执行删除、更新时有效
}
