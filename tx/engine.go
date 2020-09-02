package tx

import (
	"database/sql"
	"github.com/lyoshur/gorm/core"
)

// 事务引擎
type Engine struct {
	// 日志
	Logger core.Logger
	// 数据库事务
	tx *sql.Tx
}

// 初始化事务引擎
func (engine *Engine) Init(tx *sql.Tx) *Engine {
	engine.Logger = &core.PrintLogger{}
	engine.tx = tx
	return engine
}

// 设置日志
func (engine *Engine) SetLogger(logger core.Logger) *Engine {
	engine.Logger = logger
	return engine
}

// 执行SQL
func (engine *Engine) Exec(sql string, params []interface{}) (int64, int64, error) {
	// 执行SQL
	engine.Logger.PrintInfo(sql, params)
	result, err := engine.tx.Exec(sql, params...)
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, -1, err
	}
	// 最后插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, -1, err
	}
	// 受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, -1, err
	}
	return lastInsertId, rowsAffected, nil
}
