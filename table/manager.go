package table

import (
	"github.com/Lyo-Shur/gorm/core"
	"github.com/Lyo-Shur/gorm/utils"
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

// 查询方法 返回第一条记录
func (manager *Manager) Query(key string, param interface{}) (Table, error) {
	// 获取模板
	template := manager.SQLHolder.Get(key)
	table, err := manager.TemplateEngine.Query(template, param)
	if err != nil {
		return table, err
	}
	// 转换列名
	length := len(table.Key)
	for i := 0; i < length; i++ {
		// 转换列名
		key := manager.SQLMapper.Get(table.Key[i])
		if key != "" {
			table.Key[i] = key
		} else {
			table.Key[i] = utils.ToBigHump(table.Key[i])
		}
	}
	return table, nil
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
