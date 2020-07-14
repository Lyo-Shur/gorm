package _struct

import (
	"github.com/lyoshur/gorm/core"
	"reflect"
)

// 数据库ORM管理者
type Manager struct {
	// SQL执行引擎
	TemplateEngine *TemplateEngine
	// SQL语句缓存
	SQLHolder *core.SQLHolder
	// 字段映射缓存
	SQLMapper *core.SQLMapper
	// 结构体类型
	Model interface{}
}

// 获取数据库ORM管理者
func GetManager(templateEngine *TemplateEngine, xmlConfig core.XmlConfig, i interface{}) *Manager {
	manager := Manager{}
	manager.TemplateEngine = templateEngine
	manager.SQLHolder = (&core.SQLHolder{}).Init().Copy(xmlConfig.KVS)
	// 从结构体解析映射
	manager.SQLMapper = (&core.SQLMapper{}).Init().Copy(i)
	if xmlConfig.CPS != nil {
		for i := 0; i < len(xmlConfig.CPS); i++ {
			manager.SQLMapper.Add(xmlConfig.CPS[i].Column, xmlConfig.CPS[i].Parameter)
		}
	}
	manager.Model = i
	return &manager
}

// 查询方法 返回全部记录
func (manager *Manager) Query(key string, param interface{}) ([]reflect.Value, error) {
	// 获取模板
	template := manager.SQLHolder.Get(key)
	return manager.TemplateEngine.Query(manager.Model, manager.SQLMapper, template, param)
}

// 查询方法 返回第一条记录
func (manager *Manager) QueryFirst(key string, param interface{}) (reflect.Value, error) {
	// 获取模板
	template := manager.SQLHolder.Get(key)
	return manager.TemplateEngine.QueryFirst(manager.Model, manager.SQLMapper, template, param)
}

// 总数方法 返回count数量
func (manager *Manager) Count(key string, param interface{}) (int64, error) {
	// 获取模板
	template := manager.SQLHolder.Get(key)
	return manager.TemplateEngine.Count(template, param)
}

// 执行语句 返回最后一次插入ID和受影响的行数
func (manager *Manager) Exec(key string, param interface{}) (int64, int64, error) {
	// 获取模板
	template := manager.SQLHolder.Get(key)
	return manager.TemplateEngine.Exec(template, param)
}
