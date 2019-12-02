/*
这里对client进行二次封装，简化查询操作
结构体引擎主要使用在MVC框架中
*/

package gorm

import (
	"database/sql"
	"reflect"
)

// 结构体引擎
type structEngine struct {
	engine
}

// 获取结构体引擎
func StructEngine() *structEngine {
	engine := structEngine{}
	engine.UseDefault()
	return &engine
}

// ============================= 执行SQL方法 ============================ //

// 查询方法 返回全部记录
func (engine *structEngine) Query(model interface{}, mapper *SQLMapper, query string, args []interface{}) ([]reflect.Value, error) {
	// 执行查询
	rows, err := Client.DBS[engine.clientAlias].Link.Query(query, args...)
	if err != nil {
		return nil, err
	}
	// 返回映射结果
	return rowsToValues(rows, reflect.TypeOf(model), mapper)
}

// 查询方法 返回第一条记录
func (engine *structEngine) QueryFirst(model interface{}, mapper *SQLMapper, query string, args []interface{}) (reflect.Value, error) {
	vs, err := engine.Query(model, mapper, query, args)
	if err != nil {
		tp := reflect.TypeOf(model)
		return reflect.New(tp), err
	}
	if len(vs) == 0 {
		tp := reflect.TypeOf(model)
		return reflect.New(tp), nil
	}
	return vs[0], nil
}

// ============================= 私有*辅助方法 ============================ //

// 自动映射查询结果 sql.Rows转为reflect.Value
func rowsToValues(rows *sql.Rows, reflectType reflect.Type, m *SQLMapper) ([]reflect.Value, error) {
	// 创建返回的数组
	var returnList []reflect.Value

	// 遍历查询结果
	for rows.Next() {
		// 使用反射创建对象
		bean := reflect.New(reflectType)
		v := bean.Elem()

		// 读取所有的列信息
		ss, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		// scan装载参数
		var scanParameter []interface{}
		// 遍历所有的列信息
		for i := 0; i < len(ss); i++ {
			// 当前列名
			s := ss[i]
			// 当前列对应的参数名
			param := m.Get(s)
			// 如果当前列未找到对应的参数
			if param == "" {
				// 伪造一个指针 以免出现长度不一致
				temp := ""
				scanParameter = append(scanParameter, &temp)
			} else {
				// 从真实的结构体中获取指针
				in := v.FieldByName(param).Addr().Interface()
				scanParameter = append(scanParameter, in)
			}
		}
		// 自动装载
		err = rows.Scan(scanParameter...)
		if err != nil {
			return nil, err
		}
		// 装载结果添加到返回列表
		returnList = append(returnList, bean)
	}
	return returnList, nil
}
