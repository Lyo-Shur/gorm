/*
基础查询引擎
*/

package gorm

// 查询引擎
type engine struct {
	// 数据源名称
	clientAlias string
}

// 使用默认链接
func (e *engine) UseDefault() {
	e.Use(DefaultClientAlias)
}

// 设置使用的链接
func (e *engine) Use(clientAlias string) {
	e.clientAlias = clientAlias
}

// 多数据源 统计数量
func (e *engine) Count(query string, args []interface{}) (int64, error) {
	// 执行查询
	rows, err := Client.DBS[e.clientAlias].Link.Query(query, args...)
	if err != nil {
		return -1, err
	}
	var count int64 = 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return -1, err
		}
	}
	return count, nil
}

// 多数据源 执行SQL
func (e *engine) Exec(query string, args []interface{}) (int64, int64, error) {
	// 执行SQL
	result, err := Client.DBS[e.clientAlias].Link.Exec(query, args...)
	if err != nil {
		return -1, -1, err
	}
	// 最后插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	// 受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}
	return lastInsertId, rowsAffected, nil
}
