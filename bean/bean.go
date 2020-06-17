/*
此文件主要是一个分页使用的bean
负责查询的实体，可选择组合Page结构体，来拥有分页查询功能
*/

package bean

// 不使用分页
var closePaging int64 = -1

// 分页结构体
type Page struct {
	// 当前页码
	PageNow int64
	// 记录开始位置
	Start int64
	// 查询长度
	Length int64
}

// 分页结构体初始化方法
func (page *Page) Init() {
	page.PageNow = 1
	page.Start = 0
	page.Length = 10
}

// 计算分页方法
func (page *Page) TranslatePaging() {
	if page.Start != closePaging {
		page.Start = (page.PageNow - 1) * page.Length
	}
}

// 关闭分页方法
func (page *Page) ClosePaging() {
	page.Start = closePaging
}
