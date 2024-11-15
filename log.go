package log

import (
	"encoding/json"
	"fmt"
	"time"
)

type Color int32

const (
	Error Color = iota
	Warn
	Debug
	Info
)

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
