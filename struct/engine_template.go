package _struct

import (
	"github.com/Lyo-Shur/gorm/core"
	"github.com/Lyo-Shur/gorm/generate"
	"reflect"
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
func (templateEngine *TemplateEngine) Query(model interface{}, mapper *core.SQLMapper, template string, param interface{}) ([]reflect.Value, error) {
	// 解释模板
	sql, params := generate.SQLBuilder(template, param)
	return templateEngine.engine.Query(model, mapper, sql, params)
}

// 查询方法 返回第一条记录
func (templateEngine *TemplateEngine) QueryFirst(model interface{}, mapper *core.SQLMapper, template string, param interface{}) (reflect.Value, error) {
	// 解释模板
	sql, params := generate.SQLBuilder(template, param)
	return templateEngine.engine.QueryFirst(model, mapper, sql, params)
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
