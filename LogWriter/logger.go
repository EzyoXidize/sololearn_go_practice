package main

// 声明日志写入器接口
type LogWriter interface { // 该接口可被外部使用
	Write(data interface{}) error
}

type Logger struct { // 声明日志器结构
	writerList []LogWriter
}

// 注册一个日志写入器 ,将日志写入器的接口添加到 writeList 中
func (l *Logger) RegisterWriter(writer LogWriter)  {
	l.writerList = append(l.writerList , writer)

}

// 将一个 data 类型的数据写入日志
func (l *Logger) Log(data interface{})  {
	// 遍历所有注册的写入器
	for _,writer := range l.writerList{
		// 将日志输入到每一个写入器中
		writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
