/*
此文件主要使用Go模板引擎对SQL语句进行处理
1.为了避免SQL注入风险，在Go模板引擎上封装了#{}语法，来执行参数化查询
2.为了处理代码生成器生成UPDATE语句时，逗号的处理问题，添加了 WHERE 前的逗号检查
*/

package generate

import (
	"regexp"
	"strings"
)

// 解析SQL模板与预置参数，生成具体的SQL语句，和参数化查询所需参数
func SQLBuilder(sqlTemplate string, parameter interface{}) (string, []interface{}) {
	// 获取SQL模板并执行预处理
	st := parsing(sqlTemplate)
	// 对模板进行解释
	s := GetBuilder(st).Execute(parameter)
	// 提取模板中的语句以及参数
	sql, is := extractParameters(s)
	// 清洗SQL
	sql = strings.Trim(sql, "")
	sql = strings.Replace(sql, "\n", " ", -1)
	sql = strings.Replace(sql, "\t", " ", -1)
	l := 0
	for l != len(sql) {
		l = len(sql)
		sql = strings.Replace(sql, "  ", " ", -1)
	}
	// 处理 UPDATE
	return dealWithUpdate(sql), is
}

// 使用正则表达式对字符串模板中的#{}进行替换
// 替换为<![Parameter[{{" + str + "}}]]>
// 方便在字符串模板解析后 提取参数
// 参数提取后进行参数化查询
func parsing(str string) string {

	// 创建正则表达式
	exp := regexp.MustCompile(`(#{)(.*?)(})`)

	// 对原始字符串进行匹配
	result := exp.FindAllStringSubmatch(str, -1)

	// 循环修改匹配结果
	for i := 0; i < len(result); i++ {
		// 当前循环项
		item := result[i]
		// 原始字符串
		o := item[0]
		// 新字符串
		n := "<![Parameter[{{" + item[2] + "}}]]>"
		// 执行替换
		str = strings.Replace(str, o, n, -1)
	}
	return str
}

// 将字符串中的<![Parameter[" + str + "]]>提取出来
// 将原来字符串所在的位置替换为？
// 返回字符串替换后的结果以及提取出来的参数数组
func extractParameters(str string) (string, []interface{}) {

	// 创建正则表达式
	exp2 := regexp.MustCompile(`(<!\[Parameter\[)((.|\s|\S)*?)(\]\]>)`)

	// 对原始字符串进行匹配
	result2 := exp2.FindAllStringSubmatch(str, -1)

	// 匹配的长度
	l := len(result2)

	// 参数缓存切片
	p := make([]interface{}, l)
	for i := 0; i < l; i++ {
		// 当前循环项
		item := result2[i]

		// 原始字符串
		o := item[0]

		// 真实参数
		p[i] = item[2]

		// 替换原始字符串为？方便参数化查询
		str = strings.Replace(str, o, "?", -1)
	}
	return str, p
}

// 处理update语句中可能的错误逗号
func dealWithUpdate(sql string) string {
	// sql 语句转大写
	uper := strings.ToUpper(sql)
	// 如果其中包含 UPDATE 关键字 即是更新语句的话
	// 并且包含 WHERE 关键字
	if strings.Contains(uper, "UPDATE") && strings.Contains(uper, "WHERE") {
		// 从 WHERE 的位置向前查找
		upers := []rune(uper)
		index := strings.Index(uper, "WHERE") - 1
		for index > 0 {
			if upers[index] == rune(' ') || upers[index] == rune('\n') || upers[index] == rune('\r') {
				index--
				continue
			}
			if upers[index] == rune(',') {
				return string([]rune(sql)[:index]) + string([]rune(sql)[index+1:])
			}
			break
		}
	}
	return sql
}
