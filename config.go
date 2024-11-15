package log

// Log 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。
type Log interface {
	Error(msg string)
	Warn(msg string)
	Debug(msg string)
	Info(msg string)
}

type LogConfig struct {
	IsWriteFile bool
}

func NewLogInfo(status bool) *LogConfig {
	return &LogConfig{
		IsWriteFile: status,
	}
}

func (l *LogConfig) LogHelp(caller, module string) Log {
	if l.IsWriteFile {
		return NewFile(caller, module)
	}
	return NewTerminal(caller, module)
}
