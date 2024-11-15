package log

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
)

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
