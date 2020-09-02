package session

import (
	"database/sql"
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/gorm/table"
	"github.com/lyoshur/gorm/tx"
	"log"
)

// 会话工厂
type Factory struct {
	// 数据库连接
	DB *sql.DB
	// 日志
	Logger core.Logger
	// 会话
	Session Session
}

// 获取事务session
func (factory *Factory) TxSession() (TxSession, error) {
	Tx, err := factory.DB.Begin()
	if err != nil {
		return TxSession{}, err
	}
	txEngine := tx.Engine{}.Init(Tx).SetLogger(factory.Logger)
	return TxSession{
		Tx:               Tx,
		TxTemplateEngine: tx.GetTemplateEngine(txEngine),
	}, nil
}

// 会话工厂建造者
type FactoryBuilder struct {
	// 数据库连接
	DB *sql.DB
	// 日志
	Logger core.Logger
}

// 获取会话工厂建造者
func GetSessionFactoryBuilder(driverName string, dataSourceName string) *FactoryBuilder {
	fb := FactoryBuilder{}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	fb.DB = db
	fb.Logger = &core.PrintLogger{}
	return &fb
}

// 设置日志
func (fb *FactoryBuilder) SetLogger(logger core.Logger) *FactoryBuilder {
	fb.Logger = logger
	return fb
}

// 建造
func (fb *FactoryBuilder) Build() *Factory {
	tableEngine := table.Engine{}
	tableEngine.Init(fb.DB)
	tableEngine.SetLogger(fb.Logger)

	structEngine := _struct.Engine{}
	structEngine.Init(fb.DB)
	structEngine.SetLogger(fb.Logger)

	return &Factory{
		DB:     fb.DB,
		Logger: fb.Logger,
		Session: Session{
			TableTemplateEngine:  table.GetTemplateEngine(&tableEngine),
			StructTemplateEngine: _struct.GetTemplateEngine(&structEngine),
		},
	}
}
