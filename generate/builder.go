package generate

import (
	"bytes"
	"github.com/Lyo-Shur/gorm/tool"
	"log"
	"strings"
	"text/template"
)

// 生成器
type Builder struct {
	// 模板
	template *template.Template
}

// 获取模板生成器
func GetBuilder(tpl string) *Builder {
	builder := Builder{}
	// 创建模板
	builder.template = template.New("template")
	builder.init()
	// 解析模板
	var err error
	builder.template, err = builder.template.Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	return &builder
}

// 执行模板
func (builder *Builder) Execute(data interface{}) string {
	buf := bytes.NewBufferString("")
	// 执行模板
	err := builder.template.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

// 注册默认方法
func (builder *Builder) init() {
	builder.template.Funcs(template.FuncMap{
		"BigHump":   tool.ToBigHump,
		"SmallHump": tool.ToSmallHump,
		"Underline": tool.ToUnderline,
		"ClearType": func(Type string) string {
			return strings.Split(Type, "(")[0]
		},
	})
}
