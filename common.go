package log

type LogMessage struct {
	Ts     string `json:"ts"`     // 时间戳
	Caller string `json:"caller"` // 调用者
	Module string `json:"module"` // 模块
	Msg    string `json:"msg"`    // 消息
}

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
