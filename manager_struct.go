package gorm

import (
	"log"
	"reflect"
)

// 数据库ORM管理者
type StructManager struct {
	// SQL语句缓存
	SQLHolder *SQLHolder
	// 字段映射缓存
	SQLMapper *SQLMapper
	// SQL执行引擎
	Engine *structEngine

	// 结构体类型
	Model interface{}
}

// 获取数据库ORM管理者
func GetStructManager(xmlConfig xmlConfig, i interface{}) *StructManager {
	manager := StructManager{}
	manager.Model = i
	manager.SQLHolder = (&SQLHolder{}).Init().Copy(xmlConfig.KVS)
	manager.SQLMapper = (&SQLMapper{}).Init().Copy(i)
	if xmlConfig.CPS != nil {
		for i := 0; i < len(xmlConfig.CPS); i++ {
			manager.SQLMapper.Add(xmlConfig.CPS[i].Column, xmlConfig.CPS[i].Parameter)
		}
	}
	manager.Engine = StructEngine()
	return &manager
}

// 查询方法 返回全部记录
func (m *StructManager) QueryFirst(key string, v interface{}) (reflect.Value, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("StructManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.QueryFirst(m.Model, m.SQLMapper, sql, ps)
}

// 查询方法 返回第一条记录
func (m *StructManager) Query(key string, v interface{}) ([]reflect.Value, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("StructManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.Query(m.Model, m.SQLMapper, sql, ps)
}

// 总数方法 返回count数量
func (m *StructManager) Count(key string, v interface{}) (int64, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("StructManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.Count(sql, ps)
}

// 执行语句 返回最后一次插入ID和受影响的行数
func (m *StructManager) Exec(key string, v interface{}) (int64, int64, error) {
	template := m.SQLHolder.Get(key)
	sql, ps := SQLBuilder(template, v)
	// 打印当前执行的SQL
	log.Println("StructManager -> 当前执行的SQL:"+sql, ps)
	return m.Engine.Exec(sql, ps)
}
