package gorm

import (
	"github.com/Lyo-Shur/gorm/core"
	"github.com/Lyo-Shur/gorm/generate"
	"github.com/Lyo-Shur/gorm/struct"
	"github.com/Lyo-Shur/gorm/table"
	"github.com/Lyo-Shur/gorm/table/mvc"
	"github.com/Lyo-Shur/gorm/utils"
)

// 查询引擎
type Gorm struct{}

// 获取模板生成器
func (g Gorm) GetBuilder(tpl string) *generate.Builder {
	return generate.GetBuilder(tpl)
}

// 解析SQL模板 获取SQL和参数
func (g Gorm) SQLBuilder(sqlTemplate string, parameter interface{}) (string, []interface{}) {
	return generate.SQLBuilder(sqlTemplate, parameter)
}

// 获取基础引擎
func (g Gorm) GetEngine() *core.Engine {
	return &core.Engine{}
}

// 获取结构体引擎
func (g Gorm) GetStructEngine() *_struct.Engine {
	return &_struct.Engine{}
}

// 获取结构体模板引擎
func (g Gorm) GetStructTemplateEngine(structEngine *_struct.Engine) *_struct.TemplateEngine {
	return _struct.GetTemplateEngine(structEngine)
}

// 获取结构体Manager
func (g Gorm) GetStructManager(structTemplateEngine *_struct.TemplateEngine, xmlConfig core.XmlConfig, i interface{}) *_struct.Manager {
	return _struct.GetManager(structTemplateEngine, xmlConfig, i)
}

// 获取结构体Manager
func (g Gorm) GetStructManagerByString(structTemplateEngine *_struct.TemplateEngine, xmlConfigString string, i interface{}) *_struct.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return _struct.GetManager(structTemplateEngine, xmlConfig, i)
}

// 获取Table引擎
func (g Gorm) GetTableEngine() *table.Engine {
	return &table.Engine{}
}

// 获取Table模板引擎
func (g Gorm) GetTableTemplateEngine(tableEngine *table.Engine) *table.TemplateEngine {
	return table.GetTemplateEngine(tableEngine)
}

// 获取Table Manager
func (g Gorm) GetTableManager(tableTemplateEngine *table.TemplateEngine, xmlConfig core.XmlConfig) *table.Manager {
	return table.GetManager(tableTemplateEngine, xmlConfig)
}

// 获取Table Manager
func (g Gorm) GetTableManagerByString(tableTemplateEngine *table.TemplateEngine, xmlConfigString string) *table.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return table.GetManager(tableTemplateEngine, xmlConfig)
}

// 获取DAO
func (g Gorm) GetDAO(tableManager *table.Manager) mvc.DAO {
	return mvc.GetDAO(tableManager)
}

// 获取Service
func (g Gorm) GetService(dao mvc.DAO) mvc.Service {
	return mvc.GetService(dao)
}

// 字段名转大驼峰
func (g Gorm) ToBigHump(s string) string {
	return utils.ToBigHump(s)
}

// 字段名转小驼峰
func (g Gorm) ToSmallHump(s string) string {
	return utils.ToSmallHump(s)
}

// 字段名转下划线方法(默认去掉首字符下划线)
func (g Gorm) ToUnderline(s string) string {
	return utils.ToUnderline(s)
}

// 获取SQLHolder
func (g Gorm) GetSQLHolder() *core.SQLHolder {
	return (&core.SQLHolder{}).Init()
}

// 获取SQLMapper
func (g Gorm) GetSQLMapper() *core.SQLMapper {
	return (&core.SQLMapper{}).Init()
}

// 获取XmlConfig
func (g Gorm) GetXmlConfig(str string) core.XmlConfig {
	return core.GetXmlConfig(str)
}
