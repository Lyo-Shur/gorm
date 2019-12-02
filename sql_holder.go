/*
SQLHolder内部使用map维持了一个SQL-ID和具体SQL的映射，主要用于缓存SQL语句
数据来源:
	1.来自xml_config解析的xml配置文件
	2.来自代码生成器生成的SQL语句
*/

package gorm

type SQLHolder struct {
	// 映射关系
	m map[string]string
}

// 初始化SQLHolder
func (sqlHolder *SQLHolder) Init() *SQLHolder {
	sqlHolder.m = make(map[string]string)
	return sqlHolder
}

// 获取SQL语句
func (sqlHolder *SQLHolder) Get(key string) string {
	return sqlHolder.m[key]
}

// 添加单条映射
func (sqlHolder *SQLHolder) Add(key, value string) *SQLHolder {
	sqlHolder.m[key] = value
	return sqlHolder
}

// 从kv数组复制映射关系
// 一般kv数组来自xml_config
func (sqlHolder *SQLHolder) Copy(kvs []kv) *SQLHolder {
	for i := 0; i < len(kvs); i++ {
		sqlHolder.m[kvs[i].Key] = kvs[i].Value
	}
	return sqlHolder
}
