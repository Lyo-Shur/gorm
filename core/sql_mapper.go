/*
SQLMapper内部使用map维持了一个数据库键和结构体键的映射，主要用于将查询结果按映射反射进结构体
数据来源:
	1.来自xml_config解析的xml配置文件
	2.来自手动添加的映射关系
*/

package core

import (
	"github.com/Lyo-Shur/gorm/utils"
	"reflect"
)

// 数据库列名与结构体字段名 映射器
// map的键是下划线，也就是数据库列名
// map的键是大驼峰，也就是结构体字段名
type SQLMapper struct {
	// 映射关系
	m map[string]string
}

// 初始化SQLHolder
func (sqlMapper *SQLMapper) Init() *SQLMapper {
	sqlMapper.m = make(map[string]string)
	return sqlMapper
}

// 获取数据库列名对应的结构体字段名
func (sqlMapper *SQLMapper) Get(key string) string {
	return sqlMapper.m[key]
}

// 添加自定义映射规则
func (sqlMapper *SQLMapper) Add(key, value string) *SQLMapper {
	sqlMapper.m[key] = value
	return sqlMapper
}

// 获取映射器
func (sqlMapper *SQLMapper) Copy(i interface{}) *SQLMapper {
	// 解析字段填充映射
	addMapper(reflect.TypeOf(i), sqlMapper.m)
	return sqlMapper
}

// ============================ 以下为私有*辅助方法 ========================= //

// 根据结构体类型填充映射
func addMapper(t reflect.Type, m map[string]string) {
	// 读取字段个数
	num := t.NumField()
	// 遍历字段
	for j := 0; j < num; j++ {
		// 获取当前字段
		structFileId := t.Field(j)
		// 得到当前字段名
		v := structFileId.Name
		// 转化为下划线风格当作键
		k := utils.ToUnderline(v)
		// 保存进MAP
		m[k] = v
		// 递归检查子结构体 如果当前字段是一个结构体，递归调用
		tp := structFileId.Type
		if tp.Kind() == reflect.Struct {
			addMapper(tp, m)
		}
	}
}
