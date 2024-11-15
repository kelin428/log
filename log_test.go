package log

import "testing"

func TestLog(t *testing.T) {
	logConfig := NewLogInfo(true)
	log := logConfig.LogHelp("log/log.go:30", "server/LocalHttpRequestFilter")
	log.Info("test")
}
