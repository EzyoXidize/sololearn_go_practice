package main


import (
	"fmt"
	"os"
)

// 声明 consoleWriter 结构, 以实现命令行写入器
type consoleWriter struct {
}

// 实现LogWriter 的 write() 方法实现了日志写入接口 LogWriter 的 write() 方法
func (f *consoleWriter) Write(data interface{}) error {
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	_,err := os.Stdout.Write([]byte(str))

	return err
}

// 创建命令行写入器实例
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
