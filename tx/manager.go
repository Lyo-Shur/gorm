package tx

import (
	"github.com/lyoshur/gorm/core"
)

// 数据库ORM管理者
type Manager struct {
	// SQL执行引擎
	TemplateEngine *TemplateEngine
	// SQL语句缓存
	SQLHolder *core.SQLHolder
	// 字段映射缓存
	SQLMapper *core.SQLMapper
}

// 获取数据库ORM管理者
func GetManager(TemplateEngine *TemplateEngine, xmlConfig core.XmlConfig) *Manager {
	manager := Manager{}
	manager.TemplateEngine = TemplateEngine
	manager.SQLHolder = (&core.SQLHolder{}).Init().Copy(xmlConfig.KVS)
	manager.SQLMapper = (&core.SQLMapper{}).Init()
	if xmlConfig.CPS != nil {
		for i := 0; i < len(xmlConfig.CPS); i++ {
			manager.SQLMapper.Add(xmlConfig.CPS[i].Column, xmlConfig.CPS[i].Parameter)
		}
	}
	return &manager
}

// 执行语句 返回最后一次插入ID和受影响的行数
func (manager *Manager) Exec(key string, param interface{}) (int64, int64, error) {
	// 获取模板
	template := manager.SQLHolder.Get(key)
	return manager.TemplateEngine.Exec(template, param)
}
