/*
此文件依赖于github.com/go-sql-driver/mysql
主要使用go-sql-driver初始化了mysql连接，并保持住，方便进行数据查询
*/

package gorm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Client client

func init() {
	Client.DBS = make(map[string]db)
}

// 封装go-sql-driver数据库连接
// 提供多数据库链接支持
type client struct {
	DBS map[string]db
}
type db struct {
	// 数据库链接
	Link *sql.DB
	// 原始链接字符串
	Address string
}

// 定义默认数据库链接为default
const DefaultClientAlias = "main"

// 初始化数据库链接
func (client *client) Init(dataSource string) {
	client.MultiInit(DefaultClientAlias, dataSource)
}

// 初始化数据库链接
func (client *client) MultiInit(clientAlias, dataSource string) {
	link, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	db := db{}
	db.Link = link
	db.Address = dataSource
	Client.DBS[clientAlias] = db
}
