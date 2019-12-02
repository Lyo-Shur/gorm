package info

import (
	"gorm"
	"log"
	"strings"
)

// 表信息XML
const tableXML = `
<xml>
    <mapper column="TABLE_NAME" parameter="Name"/>
    <mapper column="ENGINE" parameter="Engine"/>
    <mapper column="TABLE_COLLATION" parameter="Collation"/>
    <mapper column="TABLE_COMMENT" parameter="Comment"/>
	<mapper column="AUTO_INCREMENT" parameter="AutoIncrement"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                TABLE_NAME, ENGINE, TABLE_COLLATION, TABLE_COMMENT, IFNULL(AUTO_INCREMENT, -1)
            FROM
                information_schema.TABLES
            WHERE
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_TYPE = 'BASE TABLE'
        </value>
    </sql>
</xml>
`

// 列信息XML
const columnXML = `
<xml>
    <mapper column="ORDINAL_POSITION" parameter="Number"/>
    <mapper column="COLUMN_NAME" parameter="Name"/>
    <mapper column="COLUMN_TYPE" parameter="Type"/>
    <mapper column="IS_NULLABLE" parameter="NullAble"/>
    <mapper column="COLUMN_DEFAULT" parameter="Defaule"/>
    <mapper column="COLUMN_COMMENT" parameter="Comment"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                ORDINAL_POSITION,
                COLUMN_NAME,
                COLUMN_TYPE,
                IS_NULLABLE,
                IFNULL(COLUMN_DEFAULT, ''),
                COLUMN_COMMENT
            FROM
                information_schema.COLUMNS
            WHERE
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_NAME = #{.TableName}
        </value>
    </sql>
</xml>
`

// 索引XML
const indexXML = `
<xml>
    <mapper column="INDEX_NAME" parameter="Name"/>
    <mapper column="COLUMN_NAME" parameter="ColumnName"/>
    <mapper column="NON_UNIQUE" parameter="Unique"/>
    <mapper column="INDEX_TYPE" parameter="Type"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                INDEX_NAME,
                COLUMN_NAME,
                NON_UNIQUE,
                INDEX_TYPE
            FROM
                information_schema.STATISTICS
            WHERE
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_NAME = #{.TableName}
        </value>
    </sql>
</xml>
`

// 外键XML
const keyXML = `
<xml>
    <mapper column="COLUMN_NAME" parameter="ColumnName"/>
    <mapper column="REFERENCED_TABLE_NAME" parameter="RelyTable"/>
    <mapper column="REFERENCED_COLUMN_NAME" parameter="RelyColumnName"/>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                COLUMN_NAME, REFERENCED_TABLE_NAME, REFERENCED_COLUMN_NAME
            FROM
                information_schema.KEY_COLUMN_USAGE
            WHERE
                CONSTRAINT_NAME != 'PRIMARY' AND
                TABLE_SCHEMA = REFERENCED_TABLE_SCHEMA AND
                TABLE_SCHEMA = #{.DataBaseName} AND TABLE_NAME = #{.TableName}
        </value>
    </sql>
</xml>
`

// 获取数据库相关信息
func GetDataBase() DataBase {
	return GetMultiDataBase(gorm.DefaultClientAlias)
}

// 获取数据库相关信息
func GetMultiDataBase(clientAlias string) DataBase {
	db := DataBase{}
	db.Alias = clientAlias
	// 解析数据库名
	address := gorm.Client.DBS[clientAlias].Address
	db.Name = strings.Split(strings.Split(address, "/")[1], "?")[0]
	// 查询数据
	db.SetTables(getTables(clientAlias, db.Name))
	for i := 0; i < len(db.Tables); i++ {
		db.Tables[i].SetColumns(getColumns(clientAlias, db.Name, db.Tables[i].Name))
		db.Tables[i].Indexs = getIndexs(clientAlias, db.Name, db.Tables[i].Name)
		db.Tables[i].Keys = getKeys(clientAlias, db.Name, db.Tables[i].Name)
	}
	return db
}

// 获取表相关信息
func getTables(clientAlias, dataBaseName string) []table {
	tableManager := gorm.GetStructManager(gorm.XmlConfig(tableXML), table{})
	tableManager.Engine.Use(clientAlias)
	vs, err := tableManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]table, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*table)
	}
	return list
}

// 获取列相关信息
func getColumns(clientAlias, dataBaseName, tableName string) []column {
	columnManager := gorm.GetStructManager(gorm.XmlConfig(columnXML), column{})
	columnManager.Engine.Use(clientAlias)
	vs, err := columnManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
		"TableName":    tableName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]column, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*column)
	}
	return list
}

// 获取索引相关信息
func getIndexs(clientAlias, dataBaseName, tableName string) []index {
	indexManager := gorm.GetStructManager(gorm.XmlConfig(indexXML), index{})
	indexManager.Engine.Use(clientAlias)
	vs, err := indexManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
		"TableName":    tableName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]index, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*index)
	}
	return list
}

// 获取外键相关信息
func getKeys(clientAlias, dataBaseName, tableName string) []key {
	keyManager := gorm.GetStructManager(gorm.XmlConfig(keyXML), key{})
	keyManager.Engine.Use(clientAlias)
	vs, err := keyManager.Query("GetList", map[string]interface{}{
		"DataBaseName": dataBaseName,
		"TableName":    tableName,
	})
	if err != nil {
		log.Println(err)
	}
	// 转换列表
	l := len(vs)
	list := make([]key, l)
	for i := 0; i < l; i++ {
		list[i] = *vs[i].Interface().(*key)
	}
	return list
}
