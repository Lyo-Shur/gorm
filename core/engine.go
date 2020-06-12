/*
基础查询引擎
提供初始化数据库连接、查询数量、执行SQL方法
*/

package core

import (
	"database/sql"
	"log"
)

// 查询引擎
type Engine struct {
	// 数据库连接
	DB *sql.DB
}

// 初始化数据库连接
func (engine *Engine) InitDB(driverName, dataSourceName string) *Engine {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	engine.DB = db
	return engine
}

// 统计数量
func (engine *Engine) Count(sql string, params []interface{}) (int64, error) {
	// 执行查询
	rows, err := engine.DB.Query(sql, params...)
	if err != nil {
		return -1, err
	}
	var count int64 = 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return -1, err
		}
	}
	return count, nil
}

// 执行SQL
func (engine *Engine) Exec(sql string, params []interface{}) (int64, int64, error) {
	// 执行SQL
	result, err := engine.DB.Exec(sql, params...)
	if err != nil {
		return -1, -1, err
	}
	// 最后插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	// 受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}
	return lastInsertId, rowsAffected, nil
}
