package log

import "github.com/sqjian/toolkit/log"

type Logger interface {
	Debugf(format string, params ...interface{})
	Infof(format string, params ...interface{})
	Warnf(format string, params ...interface{})
	Errorf(format string, params ...interface{})
}

const (
	file       = "log/bingo.log"
	maxSize    = 1
	maxBackUps = 1
	maxAge     = 1
	level      = "debug"
	console    = false
)

var defLogger Logger

func GetDefLogger() Logger {
	return defLogger
}

func init() {
	zapLog, zapLogErr := log.GenZapLog(
		file,
		maxSize,
		maxBackUps,
		maxAge,
		level,
		console,
	)
	if zapLogErr != nil {
		panic(zapLogErr)
	}
	defLogger = zapLog
}
