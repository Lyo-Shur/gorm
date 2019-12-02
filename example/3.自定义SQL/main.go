package main

import "fmt"

// 当SQL结构复杂时，支持手写SQL
// 结构使用XML 格式为 最外层为xml
// mapper 当数据库字段与程序中MAP或结构体字段名不一致时使用
// * column 数据库字段名 parameter 结构体字段名
// sql 每条SQL语句对应一个sql标签
// * key 可以理解成这条sql的别名
// * value 具体的sql
// value中可以使用完整的GO模板语法，为了方便参数化查询，还添加了#{}标签。建议使用#{}。

func main() {
	sql := `
		<xml>
			<mapper column="abc" parameter="Abc"></mapper>
    		<sql>
        		<key>GetList</key>
        		<value>
            		SELECT
                		id, title
            		FROM
                		t_test
        		</value>
    		</sql>
			<sql>
        		<key>GetCount</key>
        		<value>
            		SELECT
                		COUNT(1)
            		FROM
                		t_test
        		</value>
    		</sql>
			<sql>
        		<key>Delete</key>
        		<value>
            		Delete
            		FROM
                		t_test
					WHERE id = 1
        		</value>
    		</sql>
		</xml>
	`
	fmt.Println(sql)
}
