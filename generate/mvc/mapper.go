package mvc

import (
	"github.com/Lyo-Shur/gorm/generate"
	"github.com/Lyo-Shur/gorm/info"
)

const mapper = `
{{$ColumnNumber := (len .Table.Columns)}}
{{$key := ""}}{{range .Table.Indexs}}{{if eq .Name "PRIMARY"}}{{$key = .ColumnName}}{{end}}{{end}}
<xml>
    <sql>
        <key>GetList</key>
        <value>
            SELECT
                {{range $i, $v := .Table.Columns}}{{.Name}}{{if ne $v.Number $ColumnNumber}}, {{end}}{{end}}
            FROM
                {{.Table.Name}}
            WHERE
	{{range .Table.Columns}}
		{{if ne .Name $key}}
			{{if eq (ClearType .Type) "varchar"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
			{{if eq (ClearType .Type) "int"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}}
				{{$.LFlower}}{if ne .{{BigHump .Name}} 0}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
			{{if eq (ClearType .Type) "datetime"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}}
				{{$.LFlower}}{if ne (.{{BigHump .Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}.Format "2006-01-02 15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
		{{end}}
	{{end}}
            1 = 1
            {{$.LFlower}}{if .Start}{{$.RFlower}}{{$.LFlower}}{if ne .Start -1}{{$.RFlower}} LIMIT #{.Start},#{.Length} {{$.LFlower}}{end}{{$.RFlower}}{{$.LFlower}}{end}{{$.RFlower}}
        </value>
    </sql>
	<sql>
        <key>GetCount</key>
        <value>
            SELECT
                COUNT(1)
            FROM
                {{.Table.Name}}
            WHERE
	{{range .Table.Columns}}
		{{if ne .Name $key}}
			{{if eq (ClearType .Type) "varchar"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
			{{if eq (ClearType .Type) "int"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}}
				{{$.LFlower}}{if ne .{{BigHump .Name}} 0}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
			{{if eq (ClearType .Type) "datetime"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}}
				{{$.LFlower}}{if ne (.{{BigHump .Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}.Format "2006-01-02 15:04:05"} AND {{$.LFlower}}{end}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
		{{end}}
	{{end}}
            1 = 1
        </value>
    </sql>
	<sql>
        <key>GetModel</key>
        <value>
            SELECT
                {{range $i, $v := .Table.Columns}}{{.Name}}{{if ne $v.Number $ColumnNumber}}, {{end}}{{end}}
            FROM
                {{.Table.Name}}
            WHERE
				{{$key}} = #{.}
        </value>
    </sql>
	<sql>
        <key>Update</key>
        <value>
            UPDATE {{.Table.Name}} SET
	{{range .Table.Columns}}
		{{if ne .Name $key}}
			{{if eq (ClearType .Type) "varchar"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
			{{if eq (ClearType .Type) "int"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}}
				{{$.LFlower}}{if ne .{{BigHump .Name}} 0}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}}, {{$.LFlower}}{end}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
			{{if eq (ClearType .Type) "datetime"}}
				{{$.LFlower}}{if .{{BigHump .Name}}}{{$.RFlower}}
				{{$.LFlower}}{if ne (.{{BigHump .Name}}.Format "2006-01-02") "0001-01-01"}{{$.RFlower}} {{.Name}} = #{.{{BigHump .Name}}.Format "2006-01-02 15:04:05"}, {{$.LFlower}}{end}{{$.RFlower}}
				{{$.LFlower}}{end}{{$.RFlower}}
			{{end}}
		{{end}}
	{{end}}
            WHERE {{$key}} = #{.{{BigHump $key}}}
        </value>
    </sql>
	<sql>
        <key>Insert</key>
        <value>
            INSERT INTO {{.Table.Name}}
			(
		{{range $i, $v := .Table.Columns}}
			{{if or (ne .Name $key) (eq $.Table.AutoIncrement -1)}}
				{{.Name}}{{if ne $v.Number $ColumnNumber}}, {{end}}
			{{end}}
		{{end}}
			)
            VALUES
			(
	{{range $i, $v := .Table.Columns}}
		{{if or (ne .Name $key) (eq $.Table.AutoIncrement -1)}}
			{{if eq (ClearType .Type) "datetime"}}
				#{.{{BigHump .Name}}.Format "2006-01-02 15:04:05"}{{if ne $v.Number $ColumnNumber}}, {{end}}
			{{else}}
				#{.{{BigHump .Name}}}{{if ne $v.Number $ColumnNumber}}, {{end}}
			{{end}}
		{{end}}
	{{end}}
			)
        </value>
    </sql>
    <sql>
        <key>Delete</key>
        <value>
            DELETE FROM {{.Table.Name}} WHERE {{$key}} = #{.}
        </value>
    </sql>
</xml>
`

func init() {
	// 初始化生成器
	builder = generate.GetBuilder(mapper)
}

// 生成器
var builder *generate.Builder

// 获取MapperXML
func GetMapperXML(db info.DataBase, tableName string) string {
	return builder.Execute(map[string]interface{}{
		"LFlower":      "{",
		"RFlower":      "}",
		"DataBaseName": db.Name,
		"Table":        db.GetTable(tableName),
	})
}
