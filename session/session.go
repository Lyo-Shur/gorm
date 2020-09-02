package session

import (
	"github.com/lyoshur/gorm/core"
	"github.com/lyoshur/gorm/struct"
	"github.com/lyoshur/gorm/table"
)

// 普通会话
//noinspection SpellCheckingInspection
type Session struct {
	// Table模板引擎
	TableTemplateEngine *table.TemplateEngine
	// Struct模板引擎
	StructTemplateEngine *_struct.TemplateEngine
}

// 获取结构体Manager
//noinspection SpellCheckingInspection
func (session *Session) GetStructManager(xmlConfig core.XmlConfig, i interface{}) *_struct.Manager {
	return _struct.GetManager(session.StructTemplateEngine, xmlConfig, i)
}

// 获取结构体Manager
//noinspection SpellCheckingInspection
func (session *Session) GetStructManagerByString(xmlConfigString string, i interface{}) *_struct.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return _struct.GetManager(session.StructTemplateEngine, xmlConfig, i)
}

// 获取Table Manager
func (session *Session) GetTableManager(xmlConfig core.XmlConfig) *table.Manager {
	return table.GetManager(session.TableTemplateEngine, xmlConfig)
}

// 获取Table Manager
func (session *Session) GetTableManagerByString(xmlConfigString string) *table.Manager {
	xmlConfig := core.GetXmlConfig(xmlConfigString)
	return table.GetManager(session.TableTemplateEngine, xmlConfig)
}
