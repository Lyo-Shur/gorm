package table

import (
	"github.com/lyoshur/gorm/generate"
)

// 结构体模板引擎
type TemplateEngine struct {
	// 结构体引擎
	engine *Engine
}

// 获取结构体模板引擎
func GetTemplateEngine(engine *Engine) *TemplateEngine {
	templateEngine := TemplateEngine{}
	templateEngine.engine = engine
	return &templateEngine
}

// 查询方法 返回全部记录
func (templateEngine *TemplateEngine) Query(template string, param interface{}) (Table, error) {
	// 解释模板
	sql, params := generate.SQLBuilder(template, param)
	return templateEngine.engine.Query(sql, params)
}

// 统计数量
func (templateEngine *TemplateEngine) Count(template string, param interface{}) (int64, error) {
	// 解释模板
	sql, params := generate.SQLBuilder(template, param)
	return templateEngine.engine.Count(sql, params)
}

// 执行SQL
func (templateEngine *TemplateEngine) Exec(template string, param interface{}) (int64, int64, error) {
	// 解释模板
	sql, params := generate.SQLBuilder(template, param)
	return templateEngine.engine.Exec(sql, params)
}
