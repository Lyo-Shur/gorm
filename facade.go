package gorm

import (
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/generate"
	"github.com/lyoshur/gorm/session"
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/gorm/table"
	"github.com/lyoshur/gorm/tx"
	"github.com/lyoshur/gorm/utils"
)

// 生成器
type Builder = generate.Builder

// 获取模板生成器
//noinspection GoUnusedExportedFunction
func GetBuilder(tpl string) *generate.Builder {
	return generate.GetBuilder(tpl)
}

// 解析SQL模板 获取SQL和参数
//noinspection GoUnusedExportedFunction
func SQLBuilder(sqlTemplate string, parameter interface{}) (string, []interface{}) {
	return generate.SQLBuilder(sqlTemplate, parameter)
}

// 引擎
type Engine = core.Engine

//noinspection SpellCheckingInspection
type StructEngine = _struct.Engine
type TableEngine = table.Engine
type TxEngine = tx.Engine

// 获取结构体模板引擎
//noinspection SpellCheckingInspection
func GetStructTemplateEngine(structEngine *_struct.Engine) *_struct.TemplateEngine {
	return _struct.GetTemplateEngine(structEngine)
}

// 获取结构体Manager
//noinspection GoUnusedExportedFunction, SpellCheckingInspection
func GetStructManager(structTemplateEngine *_struct.TemplateEngine, xmlConfig core.XmlConfig, i interface{}) *_struct.Manager {
	return _struct.GetManager(structTemplateEngine, xmlConfig, i)
}

// 获取结构体Manager
//noinspection GoUnusedExportedFunction, SpellCheckingInspection
func GetStructManagerByString(structTemplateEngine *_struct.TemplateEngine, xmlConfigString string, i interface{}) *_struct.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return _struct.GetManager(structTemplateEngine, xmlConfig, i)
}

// 获取Table模板引擎
func GetTableTemplateEngine(tableEngine *table.Engine) *table.TemplateEngine {
	return table.GetTemplateEngine(tableEngine)
}

// 获取Table Manager
//noinspection GoUnusedExportedFunction
func GetTableManager(tableTemplateEngine *table.TemplateEngine, xmlConfig core.XmlConfig) *table.Manager {
	return table.GetManager(tableTemplateEngine, xmlConfig)
}

// 获取Table Manager
func GetTableManagerByString(tableTemplateEngine *table.TemplateEngine, xmlConfigString string) *table.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return table.GetManager(tableTemplateEngine, xmlConfig)
}

// 获取Table模板引擎
//noinspection GoUnusedExportedFunction
func GetTxTemplateEngine(txEngine *tx.Engine) *tx.TemplateEngine {
	return tx.GetTemplateEngine(txEngine)
}

// 获取Table Manager
//noinspection GoUnusedExportedFunction
func GetTxManager(txTemplateEngine *tx.TemplateEngine, xmlConfig core.XmlConfig) *tx.Manager {
	return tx.GetManager(txTemplateEngine, xmlConfig)
}

// 获取Table Manager
//noinspection GoUnusedExportedFunction
func GetTxManagerByString(txTemplateEngine *tx.TemplateEngine, xmlConfigString string) *tx.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return tx.GetManager(txTemplateEngine, xmlConfig)
}

// 字段名转大驼峰
//noinspection GoUnusedExportedFunction
func ToBigHump(s string) string {
	return utils.ToBigHump(s)
}

// 字段名转小驼峰
//noinspection GoUnusedExportedFunction
func ToSmallHump(s string) string {
	return utils.ToSmallHump(s)
}

// 字段名转下划线方法(默认去掉首字符下划线)
//noinspection GoUnusedExportedFunction
func ToUnderline(s string) string {
	return utils.ToUnderline(s)
}

// 获取SQLHolder
type SQLHolder = core.SQLHolder

//noinspection GoUnusedExportedFunction
func GetSQLHolder() *core.SQLHolder {
	return (&core.SQLHolder{}).Init()
}

// 获取SQLMapper
type SQLMapper = core.SQLMapper

//noinspection GoUnusedExportedFunction
func GetSQLMapper() *core.SQLMapper {
	return (&core.SQLMapper{}).Init()
}

// 获取XmlConfig
//noinspection GoUnusedExportedFunction
func GetXmlConfig(str string) core.XmlConfig {
	return core.GetXmlConfig(str)
}

type Session = session.Session
type TxSession = session.TxSession
type SessionFactory = session.Factory
type SessionFactoryBuilder = session.FactoryBuilder

// 获取会话工厂建造者
//noinspection GoUnusedExportedFunction
func GetSessionFactoryBuilder(driverName string, dataSourceName string) *SessionFactoryBuilder {
	return session.GetSessionFactoryBuilder(driverName, dataSourceName)
}
