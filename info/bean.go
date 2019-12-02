package info

// 数据库信息
type DataBase struct {
	// 链接别名
	Alias string
	// 数据库名
	Name string
	// 表列表
	Tables []table
	// 表名、下标映射
	mapper map[string]int
}

// 检查表是否存在
func (db *DataBase) TableExist(tableName string) bool {
	_, ok := db.mapper[tableName]
	return ok
}

// 设置表
func (db *DataBase) SetTables(tables []table) {
	db.Tables = tables
	db.mapper = make(map[string]int)
	for i := 0; i < len(tables); i++ {
		db.mapper[tables[i].Name] = i
	}
}

// 根据表名获取表
func (db *DataBase) GetTable(tableName string) table {
	return db.Tables[db.mapper[tableName]]
}

// 表信息
type table struct {
	// 表名
	Name string
	// 引擎
	Engine string
	// 排序规则
	Collation string
	// 注释
	Comment string
	// 自增值
	AutoIncrement int64
	// 列信息列表
	Columns []column
	// 索引信息列表
	Indexs []index
	// 外键信息
	Keys []key
	// 列名、下标映射
	mapper map[string]int
}

// 设置列
func (table *table) SetColumns(columns []column) {
	table.Columns = columns
	table.mapper = make(map[string]int)
	for i := 0; i < len(columns); i++ {
		table.mapper[columns[i].Name] = i
	}
}

// 根据列名获取列
func (table *table) GetColumn(columnName string) column {
	return table.Columns[table.mapper[columnName]]
}

// 列信息
type column struct {
	// 序号
	Number int64
	// 列名
	Name string
	// 列类型
	Type string
	// 是否允许为空
	NullAble string
	// 默认值
	Defaule string
	// 注释
	Comment string
}

// 索引信息
type index struct {
	// 索引名
	Name string
	// 列名
	ColumnName string
	// 是否是唯一约束
	Unique string
	// 索引类型
	Type string
}

// 键信息
type key struct {
	// 列名account.xml
	ColumnName string
	// 依赖的表
	RelyTable string
	// 依赖的列名
	RelyColumnName string
}
