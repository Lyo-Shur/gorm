package gorm

import (
	"log"
)

// 数据库ORM管理者
type TableManager struct {
	// SQL语句缓存
	SQLHolder *SQLHolder
	// 字段映射缓存
	SQLMapper *SQLMapper
	// SQL执行引擎
	Engine *tableEngine
}

// 获取数据库ORM管理者
func GetTableManager(xmlConfig xmlConfig) *TableManager {
	manager := TableManager{}
	manager.SQLHolder = (&SQLHolder{}).Init().Copy(xmlConfig.KVS)
	manager.SQLMapper = (&SQLMapper{}).Init()
	if xmlConfig.CPS != nil {
		for i := 0; i < len(xmlConfig.CPS); i++ {
			manager.SQLMapper.Add(xmlConfig.CPS[i].Column, xmlConfig.CPS[i].Parameter)
		}
	}
	manager.Engine = TableEngine()
	return &manager
}

// 查询方法 返回第一条记录
func (m *TableManager) Query(key string, v interface{}) (Table, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("TableManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.Query(m.SQLMapper, sql, ps)
}

// 总数方法 返回count数量
func (m *TableManager) Count(key string, v interface{}) (int64, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("TableManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.Count(sql, ps)
}

// 执行语句 返回最后一次插入ID和受影响的行数
func (m *TableManager) Exec(key string, v interface{}) (int64, int64, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("TableManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.Exec(sql, ps)
}
