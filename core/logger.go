package core

import "fmt"

// 日志接口
type Logger interface {
	// 打印日志
	PrintInfo(sql string, parameters []interface{})
	// 打印异常
	PrintError(err error)
}

// 无操作日志实现
type NoLogger struct{}

func (noLogger *NoLogger) PrintInfo(sql string, parameters []interface{}) {}
func (noLogger *NoLogger) PrintError(err error)                           {}

// 普通打印日志实现
type PrintLogger struct{}

func (printLogger *PrintLogger) PrintInfo(sql string, parameters []interface{}) {
	prefix := "Currently executing SQL>>>"
	fmt.Println(prefix, sql, parameters)
}
func (printLogger *PrintLogger) PrintError(err error) {
	prefix := "SQL execution error>>>"
	fmt.Println(prefix, err)
}
