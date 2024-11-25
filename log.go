package log

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
)

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

// 提取通用的日志类型
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

// 终端日志
type Terminal struct {
	LogMessage
}

func NewTerminal(caller, module string) Log {
	return &Terminal{
		LogMessage: LogMessage{Caller: caller, Module: module},
	}
}

func (t *Terminal) Debug(msg string) {
	t.logMessage(msg, Debug)
}

func (t *Terminal) Warn(msg string) {
	t.logMessage(msg, Warn)
}

func (t *Terminal) Error(msg string) {
	t.logMessage(msg, Error)
}

func (t *Terminal) Info(msg string) {
	t.logMessage(msg, Info)
}

func (t *Terminal) logMessage(msg string, color Color) {
	t.Msg = msg
	t.Ts = time.Now().Format(time.RFC3339)
	levelStr, levelName := GetType(color)
	logInfoStr := t.formatLog(levelStr, levelName)
	fmt.Println(logInfoStr)
}

func (t *Terminal) formatLog(levelStr, levelName string) string {
	reset := "\033[0m"
	logInfo, _ := json.Marshal(t.LogMessage)
	return fmt.Sprintf("%s %s[%s]%s %s", time.Now().Format("2006-01-02 15:04:05"), levelStr, levelName, reset, string(logInfo))
}

// 文件日志
type File struct {
	LogMessage
}

func NewFile(caller, module string) Log {
	return &File{
		LogMessage: LogMessage{Caller: caller, Module: module},
	}
}

func (f *File) Debug(msg string) {
	f.logMessage(msg, Debug)
}

func (f *File) Warn(msg string) {
	f.logMessage(msg, Warn)
}

func (f *File) Error(msg string) {
	f.logMessage(msg, Error)
}

func (f *File) Info(msg string) {
	f.logMessage(msg, Info)
}

func (f *File) logMessage(msg string, color Color) {
	f.Msg = msg
	f.Ts = time.Now().Format(time.RFC3339)
	levelStr, levelName := GetType(color)
	logInfoStr := f.formatLog(levelStr, levelName)
	fmt.Printf(logInfoStr)                        // 打印到控制台
	f.FileOperation(removeColorCodes(logInfoStr)) // 保存到文件，去除颜色
}

func (f *File) formatLog(levelStr, levelName string) string {
	reset := "\033[0m"
	logInfo, _ := json.Marshal(f.LogMessage)
	return fmt.Sprintf("%s %s[%s]%s %s\n", time.Now().Format("2006-01-02 15:04:05"), levelStr, levelName, reset, string(logInfo))
}

func (f *File) FileOperation(msg string) {
	file, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("failed to open file:", err)
	}
	defer file.Close()

	// 创建一个缓冲写入器
	writer := bufio.NewWriter(file)

	// 将消息写入缓冲区
	_, err = writer.WriteString(msg)
	if err != nil {
		fmt.Println("failed to write to file:", err)
	}

	// 确保缓冲区的数据被写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("failed to flush buffer:", err)
	}
}

// 去除ANSI颜色编码
func removeColorCodes(input string) string {
	// 正则表达式匹配ANSI颜色编码
	re := regexp.MustCompile(`\033\[[0-9;]*m`)
	// 替换所有匹配的ANSI转义码为空字符串
	return re.ReplaceAllString(input, "")
}
