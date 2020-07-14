package gorm

import (
	"github.com/lyoshur/gorm/bean"
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/generate"
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/gorm/table"
	"github.com/lyoshur/gorm/table/mvc"
	"github.com/lyoshur/gorm/utils"
)

// 实体
type Page = bean.Page

// 生成器
type Builder = generate.Builder

// 获取模板生成器
func GetBuilder(tpl string) *generate.Builder {
	return generate.GetBuilder(tpl)
}

// 解析SQL模板 获取SQL和参数
func SQLBuilder(sqlTemplate string, parameter interface{}) (string, []interface{}) {
	return generate.SQLBuilder(sqlTemplate, parameter)
}

// 引擎
type Engine = core.Engine
type StructEngine = _struct.Engine
type TableEngine = table.Engine

// 获取结构体模板引擎
func GetStructTemplateEngine(structEngine *_struct.Engine) *_struct.TemplateEngine {
	return _struct.GetTemplateEngine(structEngine)
}

// 获取结构体Manager
func GetStructManager(structTemplateEngine *_struct.TemplateEngine, xmlConfig core.XmlConfig, i interface{}) *_struct.Manager {
	return _struct.GetManager(structTemplateEngine, xmlConfig, i)
}

// 获取结构体Manager
func GetStructManagerByString(structTemplateEngine *_struct.TemplateEngine, xmlConfigString string, i interface{}) *_struct.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return _struct.GetManager(structTemplateEngine, xmlConfig, i)
}

// 获取Table模板引擎
func GetTableTemplateEngine(tableEngine *table.Engine) *table.TemplateEngine {
	return table.GetTemplateEngine(tableEngine)
}

// 获取Table Manager
func GetTableManager(tableTemplateEngine *table.TemplateEngine, xmlConfig core.XmlConfig) *table.Manager {
	return table.GetManager(tableTemplateEngine, xmlConfig)
}

// 获取Table Manager
func GetTableManagerByString(tableTemplateEngine *table.TemplateEngine, xmlConfigString string) *table.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return table.GetManager(tableTemplateEngine, xmlConfig)
}

// 获取DAO
func GetDAO(tableManager *table.Manager) mvc.DAO {
	return mvc.GetDAO(tableManager)
}

// 获取Service
func GetService(dao mvc.DAO) mvc.Service {
	return mvc.GetService(dao)
}

// 字段名转大驼峰
func ToBigHump(s string) string {
	return utils.ToBigHump(s)
}

// 字段名转小驼峰
func ToSmallHump(s string) string {
	return utils.ToSmallHump(s)
}

// 字段名转下划线方法(默认去掉首字符下划线)
func ToUnderline(s string) string {
	return utils.ToUnderline(s)
}

// 获取SQLHolder
type SQLHolder = core.SQLHolder

func GetSQLHolder() *core.SQLHolder {
	return (&core.SQLHolder{}).Init()
}

// 获取SQLMapper
type SQLMapper = core.SQLMapper

func GetSQLMapper() *core.SQLMapper {
	return (&core.SQLMapper{}).Init()
}

// 获取XmlConfig
func GetXmlConfig(str string) core.XmlConfig {
	return core.GetXmlConfig(str)
}
