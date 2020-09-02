package session

import (
	"database/sql"
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/tx"
)

// 事务会话
type TxSession struct {
	// 数据库事务
	Tx *sql.Tx
	// 事务模板引擎
	TxTemplateEngine *tx.TemplateEngine
}

// 提交事务
func (session *TxSession) Commit() error {
	return session.Tx.Commit()
}

// 回滚事务
func (session *TxSession) Rollback() error {
	return session.Tx.Rollback()
}

// 获取Table Manager
//noinspection GoUnusedExportedFunction
func (session *TxSession) GetTxManager(xmlConfig core.XmlConfig) *tx.Manager {
	return tx.GetManager(session.TxTemplateEngine, xmlConfig)
}

// 获取Table Manager
//noinspection GoUnusedExportedFunction
func (session *TxSession) GetTxManagerByString(xmlConfigString string) *tx.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return tx.GetManager(session.TxTemplateEngine, xmlConfig)
}
