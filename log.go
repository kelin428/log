package log

// Log 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。
type Log interface {
	Error(msg string)
	Warn(msg string)
	Debug(msg string)
	Info(msg string)
}

type Config struct {
	IsWriteFile bool
}

func NewConfig(status bool) *Config {
	return &Config{
		IsWriteFile: status,
	}
}

func (l *Config) LogHelp(caller, module string) Log {
	if l.IsWriteFile {
		return NewFile(caller, module)
	}
	return NewTerminal(caller, module)
}

type Color int32

const (
	Error Color = iota
	Warn
	Debug
	Info
)

type LogMessage struct {
	Ts     string `json:"ts"`     // 时间戳
	Caller string `json:"caller"` // 调用者
	Module string `json:"module"` // 模块
	Msg    string `json:"msg"`    // 消息
}

// GetType 提取通用的日志类型
func GetType(color Color) (string, string) {
	var levelName, levelStr string
	switch color {
	case Debug:
		levelName = "Debug"
		levelStr = "\033[35m"
	case Info:
		levelName = "Info"
		levelStr = "\033[32m"
	case Warn:
		levelName = "Warn"
		levelStr = "\033[33m"
	case Error:
		levelName = "Error"
		levelStr = "\033[31m"
	default:
		levelName = "Unknown"
		levelStr = "\033[0m"
	}

	return levelStr, levelName
}
