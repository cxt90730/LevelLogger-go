package LevelLogger_go

import (
	"fmt"
	"log"
)

const (
	LogError = iota
	LogInfo
	LogWarning
	LogDebug
)

type LevelLogger struct {
	logger   *log.Logger
	logLevel int
}

func (l *LevelLogger) Error(v ...interface{}) {
	if l.logLevel >= LogError {
		l.PrintLevelLog("ERROR", v...)
	}
}

func (l *LevelLogger) Warning(v ...interface{}) {
	if l.logLevel >= LogWarning {
		l.PrintLevelLog("WARNING", v...)
	}
}

func (l *LevelLogger) Info(v ...interface{}) {
	if l.logLevel >= LogInfo {
		l.PrintLevelLog("INFO", v...)
	}
}

func (l *LevelLogger) Debug(v ...interface{}) {
	if l.logLevel >= LogDebug {
		l.PrintLevelLog("DEBUG", v...)
	}
}

func (l *LevelLogger) PrintLevelLog(level string, v ...interface{}) {
	printLevelLog(l.logger, level, v...)
}

func printLevelLog(logger *log.Logger, level string, v ...interface{}) {
	f := logFormat(level, v...)
	logger.Printf(f, v...)
}

func logFormat(level string, v ...interface{}) string {
	formatStr := fmt.Sprintf("[%s] ", level)
	for i := 0; i < len(v); i++ {
		formatStr += "%v "
	}
	formatStr += "\n"
	return formatStr
}
