/*
这里对client进行二次封装，简化查询操作
表结构引擎与结构体引擎相比，无需传递具体结构体，针对无结构体查询使用
*/

package gorm

import (
	"database/sql"
	"github.com/Lyo-Shur/gorm/tool"
	"reflect"
)

// MYSQL查询结果集
type Table struct {
	// 列名数组
	Key []string
	// 数据集
	Values [][]interface{}
}

// 对数据进行清洗
func (table *Table) clear() {
	for i := 0; i < len(table.Values); i++ {
		for j := 0; j < len(table.Key); j++ {
			value := table.Values[i][j]
			// 当查询结果出现[]uint8切片时，证明此字段为字符串，需要手动进行转码
			tp := reflect.TypeOf(value).Kind().String()
			if tp == "slice" {
				table.Values[i][j] = string(value.([]uint8))
			}
		}
	}
}

// 转MAP
func (table *Table) ToMap() []map[string]interface{} {
	var maps []map[string]interface{}
	// 遍历行
	for i := 0; i < len(table.Values); i++ {
		// 处理当前行
		m := make(map[string]interface{})
		for j := 0; j < len(table.Key); j++ {
			m[table.Key[j]] = table.Values[i][j]
		}
		maps = append(maps, m)
	}
	return maps
}

// 表结构引擎
type tableEngine struct {
	engine
}

// 获取表结构引擎
func TableEngine() *tableEngine {
	engine := tableEngine{}
	engine.UseDefault()
	return &engine
}

// ============================= 执行SQL方法 ============================ //

// 查询方法 返回全部记录
func (engine *tableEngine) Query(mapper *SQLMapper, query string, args []interface{}) (Table, error) {
	// 执行查询
	rows, err := Client.DBS[engine.clientAlias].Link.Query(query, args...)
	if err != nil {
		return Table{}, err
	}
	// 返回映射结果
	return rowsToTable(rows, mapper)
}

// ============================= 私有*辅助方法 ============================ //

// 自动映射查询结果 sql.Rows转为Table
func rowsToTable(rows *sql.Rows, mapper *SQLMapper) (Table, error) {
	var table Table
	for rows.Next() {
		// 读取所有的列信息
		ss, err := rows.Columns()
		if err != nil {
			return table, err
		}
		table.Key = ss

		// scan装载参数
		l := len(table.Key)
		value := make([]interface{}, l)
		valuePointer := make([]interface{}, l)
		// 遍历所有的列信息
		for i := 0; i < l; i++ {
			// 转换列名
			key := mapper.Get(table.Key[i])
			if key != "" {
				table.Key[i] = key
			} else {
				table.Key[i] = tool.ToBigHump(table.Key[i])
			}
			// 储存指针
			valuePointer[i] = &value[i]
		}

		// 自动装载
		err = rows.Scan(valuePointer...)
		if err != nil {
			return table, err
		}
		// 装载结果添加到返回列表
		table.Values = append(table.Values, value)
	}
	table.clear()
	return table, nil
}
